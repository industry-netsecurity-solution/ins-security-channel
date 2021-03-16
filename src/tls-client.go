package main

import (
	"crypto/tls"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/common"
	"github.com/industry-netsecurity-solution/ins-security-channel/config"
	"github.com/industry-netsecurity-solution/ins-security-channel/utils"
	"github.com/spf13/viper"
	"net"
	"os"
	"path"
	"strconv"
	"time"
)

var ClientConfig config.ClientConfigurations

func LoadClientConfiguration(configFile string) {
	// Set the file name of the configurations file
	//viper.AddConfigPath(configFile)
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Printf("%s", dir)

	err = viper.Unmarshal(&ClientConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading config file
	fmt.Println("Reading config file")
	fmt.Println("RemoteTlsServerIp is\t", ClientConfig.RemoteTlsServerIp)
	fmt.Println("RemoteTlsServerPort is\t\t", ClientConfig.RemoteTlsServerPort)
}

func main() {

	configPath := flag.String("c", "config.yaml", "configuration path(default:config.yaml)")
	fileType := flag.Uint("f", 0, "file type (0: normal, 1: front video, 2: rear video, 3: front collision, 4: rear collision, 5: front approach, 6: rear approach)")
	srcId := flag.String("s", "smart-gw-01", "file type (0: normal, 1: front video, 2: rear video, 3: front collision, 4: rear collision, 5: front approach, 6: rear approach)")

	flag.Parse()

	// 프로그램 인자 확인
	args := flag.Args()

	if args == nil || len(args) == 0 {
		fmt.Println("Filename Required...")
		return
	}

	if fileType == nil {
		fmt.Println("FileType Required...")
		return
	}

	// 설정 파일 읽기
	LoadClientConfiguration(*configPath)

	// TLS 서버 접속 정보
	RemoteTlsServerIp := ClientConfig.RemoteTlsServerIp
	RemoteTlsServerPort := ClientConfig.RemoteTlsServerPort

	if &RemoteTlsServerIp == nil || len(RemoteTlsServerIp) == 0 {
		fmt.Println("RemoteTlsServerIp Required...")
		return
	}

	if &RemoteTlsServerPort == nil || RemoteTlsServerPort == 0 {
		fmt.Println("RemoteTlsServerPort Required...")
		return
	}

	// TLS 서버 주소
	tlsServerURL := RemoteTlsServerIp + ":" + strconv.Itoa(RemoteTlsServerPort)
	rAddr, err := net.ResolveTCPAddr("tcp", tlsServerURL)
	if err != nil {
		fmt.Println(err)
	}

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	// 로그 기록
	evt := common.EventLog{}
	evt.SetEventGatewayType("스마트게이트웨이")
	if *fileType == 0 {
		evt.SetEventType("영상 파일")
	} else if *fileType == 1 {
		evt.SetEventType("충돌 로그")
	} else if *fileType == 2 {
		evt.SetEventType("접근 감지 로그")
	}
	evt.SetEventGatewayId(*srcId)

	for index, filename := range args {
		fmt.Println(index, filename)

		stat, err := os.Stat(filename)
		if err != nil {
			t := time.Now()
			evt.SetEventTime(fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second()))

			evt.SetEventStatus("점검")
			evt.SetEventMessage(fmt.Sprintf("%s 파일 전송 실패", filename))
			evt.SetEventContent(err.Error())
			utils.ReportLog(evt)

			panic(err)
			break
		}

		file, err := os.Open(filename)
		defer file.Close()

		// TLS 연결
		tlsConn, err := tls.Dial("tcp", rAddr.String(), conf)
		if err != nil {
			//fmt.Println(err)
			t := time.Now()
			evt.SetEventTime(fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second()))

			evt.SetEventStatus("점검")
			evt.SetEventMessage(fmt.Sprintf("%s 파일 전송 실패", filename))
			evt.SetEventContent(err.Error())
			utils.ReportLog(evt)

			panic(err)
			break
		}

		params := make(map[int][]byte)

		// source
		params[0] = []byte(*srcId)

		// file type
		params[1] = make([]byte, 2)
		binary.LittleEndian.PutUint16(params[1], uint16(*fileType))

		// file name
		params[2] = []byte(path.Base(filename))

		// file size
		params[3] = make([]byte, 4)
		binary.LittleEndian.PutUint32(params[3], uint32(stat.Size()))

		err = utils.SendFileOverTls(file, params, tlsConn)
		if err != nil {
			t := time.Now()
			evt.SetEventTime(fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second()))

			evt.SetEventStatus("점검")
			evt.SetEventMessage(fmt.Sprintf("%s 파일 전송 실패", filename))
			evt.SetEventContent(err.Error())
			utils.ReportLog(evt)

			panic(err)
			break
		}

		t := time.Now()
		evt.SetEventTime(fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second()))

		evt.SetEventStatus("정상")
		evt.SetEventMessage(fmt.Sprintf("%s 파일 전송 성공", filename))
		utils.ReportLog(evt)
	}
}
