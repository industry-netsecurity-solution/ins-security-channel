package main

import (
	"flag"
	"fmt"
	isc "github.com/industry-netsecurity-solution/ins-security-channel/isc"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

var RelayConfig isc.RelayConfigurations

var GConfigPath *string = nil
var GSrcId *string = nil

var Log *log.Logger = nil


/*
 *	parse command line
 */
func ParseOptions() int {
	GConfigPath = flag.String("c", "config.yaml", "configuration path(default:config.yaml)")
	GSrcId = flag.String("s", "vip-01-relay", "service name")

	flag.Parse()

	return 0
}

func PrintConfig() {
	// Reading config file
	for _, item := range  RelayConfig.ToString() {
		Log.Println(item)
	}
}

/**
 * config를 로딩한다.
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
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&RelayConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	if RelayConfig.LocalServerPort == 0 {
		Log.Println("LocalTlsServerPort Required...")
		return -1
	}

	if len(RelayConfig.RemoteServerIp) == 0 {
		fmt.Println("RemoteTlsServerIp Required...")
		return -1
	}

	if RelayConfig.RemoteServerPort == 0 {
		Log.Println("RemoteTlsServerPort Required...")
		return -1
	}

	if len(RelayConfig.TlsCert) == 0 {
		Log.Println("TlsCert Required...")
		return -1
	}

	if len(RelayConfig.TlsKey) == 0 {
		Log.Println("TlsKey Required...")
		return -1
	}

	PrintConfig()

	return 0
}

func callbck_event(data1 []byte, data2 []byte, err *error) {
	EventLogUrl := RelayConfig.EventLogUrl
	if len(EventLogUrl) == 0 {
		return
	}

	if err != nil {
		// 로그 기록
		evt := isc.EventLog{}
		evt.SetEventGatewayType("스마트게이트웨이")
		evt.SetEventType("데이터 브로커")
		evt.SetEventGatewayId(RelayConfig.SourceId)
		t := time.Now()
		evt.SetEventTime(fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second()))

		evt.SetEventStatus("점검")
		evt.SetEventMessage((*err).Error())
		evt.SetEventContent((*err).Error())
		isc.ReportLog(RelayConfig.EventLogUrl, evt)
	} else {
		// 로그 기록
		evt := isc.EventLog{}
		evt.SetEventGatewayType("지능형 플랫폼")
		evt.SetEventType("TLS->TCP 데이터 브로커")
		evt.SetEventGatewayId(RelayConfig.SourceId)
		t := time.Now()
		evt.SetEventTime(fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second()))

		evt.SetEventStatus("성공")
		evt.SetEventMessage("")
		evt.SetEventContent("")
		isc.ReportLog(RelayConfig.EventLogUrl, evt)
	}
}

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

	if GSrcId != nil && 0 < len(*GSrcId) {
		RelayConfig.SourceId = *GSrcId
	}

	if len(RelayConfig.SourceId) == 0 {
		Log.Println("SourceId Required...")
		return
	}

	isc.InitTLS2TLS(RelayConfig, callbck_event)
}
