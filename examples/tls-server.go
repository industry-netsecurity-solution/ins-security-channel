package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	isc "github.com/industry-netsecurity-solution/ins-security-channel/isc"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
	"time"
)

var GConfigPath *string = nil
var GSrcId *string = nil

var ServerConfig isc.ServerConfigurations

var Log *log.Logger = nil

/*
 *	parse command line
 */
func ParseOptions() int {
	GConfigPath = flag.String("c", "config.yaml", "configuration path(default:config.yaml)")
	GSrcId = flag.String("s", "vip-server", "지능형 플랫폼")

	flag.Parse()

	return 0
}

/*
 * load configuration
 */
func LoadConfiguration() int {

	if GConfigPath == nil {
		return -1
	}

	// Set the file name of the configurations file
	//viper.AddConfigPath(configFile)
	viper.SetConfigFile(*GConfigPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		Log.Printf("Error reading config file, %s\n", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		return -1
	}
	Log.Printf("%s", dir)

	err = viper.Unmarshal(&ServerConfig)
	if err != nil {
		Log.Printf("Unable to decode into struct, %v\n", err)
		return -1
	}

	// Reading config file
	for _, item := range  ServerConfig.ToString() {
		Log.Println(item)
	}

	return 0
}

func handle_error_log(err error) {
	EventLogUrl := ServerConfig.EventLogUrl
	if len(EventLogUrl) == 0 {
		return
	}

	// 로그 기록
	evt := isc.EventLog{}
	evt.SetEventGatewayType("IoT 보안 시스템")
	evt.SetEventType("IoT 보안 시스템")
	evt.SetEventGatewayId(ServerConfig.SourceId)
	t := time.Now()
	evt.SetEventTime(fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()))

	evt.SetEventStatus("점검")
	evt.SetEventMessage(err.Error())
	evt.SetEventContent(err.Error())
	isc.ReportLog(EventLogUrl, evt)

}

/**
 * 연결로 부터 데이터를 읽어 파일로 저장한다.
 */
func receiveData(conn net.Conn) error {

	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	var sourceId string
	var fileType uint16
	var fileName string
	var filesize uint32

	// Tag
	rlen, err := conn.Read(buf2)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("recv tag: %d\n", rlen)

	// Length
	rlen, err = conn.Read(buf4)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("recv length: %d\n", rlen)
	payloadLength := binary.LittleEndian.Uint32(buf4)
	fmt.Printf("payload length: %d\n", payloadLength)

	for i := 0; i < 4; i++ {
		// Tag
		rlen, err = conn.Read(buf2)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Tag
		rlen, err = conn.Read(buf4)
		if err != nil {
			fmt.Println(err)
			return err
		}

		tag := binary.LittleEndian.Uint16(buf2)
		length := binary.LittleEndian.Uint32(buf4)

		fmt.Printf("tag: %d\n", tag)
		fmt.Printf("length: %d\n", length)

		bufdata := make([]byte, length)
		// Tag
		rlen, err = conn.Read(bufdata)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if tag == 0 {
			sourceId = string(bufdata)
			fmt.Printf("Source: %s\n", sourceId)
		} else if tag == 1 {
			fileType = binary.LittleEndian.Uint16(bufdata)
			fmt.Printf("fileType: %d\n", fileType)
		} else if tag == 2 {
			fileName = string(bufdata)
			fmt.Printf("fileName: %s\n", fileName)
		} else if tag == 3 {
			filesize = binary.LittleEndian.Uint32(bufdata)
			fmt.Printf("filesize: %d\n", filesize)
		}
	}

	// Tag
	rlen, err = conn.Read(buf2)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Length
	rlen, err = conn.Read(buf4)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// tag := binary.LittleEndian.Uint16(buf2)
	// length := binary.LittleEndian.Uint32(buf4)

	var dirpath string

	if fileType == 1 {
		dirpath = "./video0/normal"
	} else if fileType == 2 {
		dirpath = "./video1/normal"
	} else if fileType == 3 {
		dirpath = "./video0/collision"
	} else if fileType == 4 {
		dirpath = "./video1/collision"
	} else if fileType == 5 {
		dirpath = "./video0/approach"
	} else if fileType == 6 {
		dirpath = "./video1/approach"
	}

	var filePath string
	if 0 < len(dirpath) {
		err = os.MkdirAll(dirpath, 0755)
		if err != nil {
			fmt.Println(err)
			return err
		}
		filePath = fmt.Sprintf("%s/%s", dirpath, fileName)
	} else {
		filePath = fileName
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	writer := bufio.NewWriter(file)

	defer func() {
		writer.Flush()
		file.Close()
	}()

	for {
		buf := make([]byte, 4096)
		rlen, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return err
		}

		//fmt.Printf("rlen = %d\n", rlen)

		res := make([]byte, rlen)
		copy(res, buf[:rlen])

		_, err = writer.Write(res)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
}

/**
 * 파일 수신 TLS 서버
 */
func main() {
	Log = log.New(os.Stdout, "", log.LstdFlags)

	// 프로그램 인자 확인
	if ParseOptions() == -1 {
		return
	}

	// 설정 파일 읽기
	if LoadConfiguration() == -1 {
		return
	}

	if ServerConfig.LocalServerPort == 0 {
		fmt.Println("LocalTlsServerPort Required...")
		return
	}

	if 0 < len(*GSrcId) {
		ServerConfig.SourceId = *GSrcId
	}

	isc.ReadyServer(&ServerConfig, receiveData)
}
