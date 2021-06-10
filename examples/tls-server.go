package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
	"time"
)

var Config ins.RelayConfigurations

var GConfigPath *string = nil
var GSrcId *string = nil

var GServiceEnableTls *bool = nil
var GServiceAddress *string = nil
var GServicePort *int = nil

var Log *log.Logger = nil

/*
 *	parse command line
 */
func ParseFlagOptions() int {

	GConfigPath = flag.String("c", "config.yaml", "configuration path(default:config.yaml)")
	GSrcId = flag.String("source", "ins-security-server", "IoT 보안 시스템")

	GServiceEnableTls = flag.Bool("service.enabletls", false, "The address of service.")
	GServiceAddress = flag.String("service.address", "0.0.0.0", "The address of service.")
	GServicePort = flag.Int("service.port", 9980, "The port of service.")

	flag.Parse()

	return 0
}

/*
 * load configuration
 */
func LoadConfiguration() int {
	if GConfigPath == nil || len(*GConfigPath) <= 0 {
		return 0
	}

	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	// Set the file name of the configurations file
	//viper.AddConfigPath(configFile)
	viper.SetConfigFile(*GConfigPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		Log.Printf("Error reading config file, %s", err.Error())
	}

	if viper.Get("Service.EnableTls") == nil {
		viper.SetDefault("Service.EnableTls", false)
	}

	if viper.Get("Service.Address") == nil {
		viper.SetDefault("Service.Address", "0.0.0.0")
	}

	if viper.Get("Service.Port") == nil {
		viper.SetDefault("Service.Port", 9980)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		Log.Printf("Unable to decode into struct, %v", err)
		return -1
	}

	return 0
}

/*
 * load configuration
 */
func UpdateConfiguration() int {

	// Service
	if ins.IsFlagPassed("service.enabletls") {
		Config.Service.EnableTls = *GServiceEnableTls
	}
	if ins.IsFlagPassed("service.address") {
		Config.Service.Address = *GServiceAddress
	} else {
		if len(Config.Service.Address) == 0 {
			Config.Service.Address = *GServiceAddress
		}
	}

	if ins.IsFlagPassed("service.port") {
		Config.Service.Port = *GServicePort
	} else {
		if Config.Service.Port == 0 {
			Config.Service.Port = *GServicePort
		}
	}

	//
	if ins.IsFlagPassed("source") {
		Config.SourceId = *GSrcId
	} else {
		if len(Config.SourceId) == 0 {
			Config.SourceId = *GSrcId
		}
	}

	// Reading config file
	Config.PrintConfigurations()

	return 0
}

func ParseOptions() int {

	// 프로그램 인자 확인
	if ParseFlagOptions() == -1 {
		return -1
	}

	// 설정 파일 읽기
	if LoadConfiguration() == -1 {
		return -1
	}

	// 설정 파일 읽기
	if UpdateConfiguration() == -1 {
		return -1
	}

	Config.PrintConfigurations()

	return 0
}

func handle_error_log(err error) {
	EventLogUrl := Config.EventLogUrl
	if len(EventLogUrl) == 0 {
		return
	}

	// 로그 기록
	evt := ins.EventLog{}
	evt.SetEventGatewayType("IoT 보안 시스템")
	evt.SetEventType("IoT 보안 시스템")
	evt.SetEventGatewayId(Config.SourceId)
	t := time.Now()
	evt.SetEventTime(fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()))

	evt.SetEventStatus("점검")
	evt.SetEventMessage(err.Error())
	evt.SetEventContent(err.Error())
	ins.ReportEvent(EventLogUrl, evt)

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

func callback(conn net.Conn, ud interface{}) error {

	for {
		data, err := ins.RecvTLV(conn, binary.LittleEndian)
		if err != nil {
			return err
		}
		fmt.Println(data)
	}

	return nil
}

/**
 * 파일 수신 서버
 */
func main() {

	Log = log.New(os.Stdout, "", log.LstdFlags)

	if ParseOptions() != 0 {
		return
	}

	ins.ReadyServer(&Config.Service, nil, callback)
}
