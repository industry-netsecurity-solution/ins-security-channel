package main

import (
	"flag"
	"fmt"
	isc "github.com/industry-netsecurity-solution/ins-security-channel/isc"
	"github.com/spf13/viper"
	"os"
	"time"
)

var RelayConfig isc.RelayConfigurations

/**
 * config를 로딩한다.
 */
func LoadRelayConfiguration(configFile string) {
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

	err = viper.Unmarshal(&RelayConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading config file
	fmt.Println("Reading config file")
	fmt.Println("TlsCert is\t", RelayConfig.TlsCert)
	fmt.Println("TlsKey is\t", RelayConfig.TlsKey)
	fmt.Println("LocalTlsServerPort is\t", RelayConfig.LocalTlsServerPort)
	fmt.Println("RemoteTlsServerIp is\t\t", RelayConfig.RemoteTlsServerIp)
	fmt.Println("RemoteTlsServerPort is\t\t", RelayConfig.RemoteTlsServerPort)
}

func callbck_handle_error(err error) {
	EventLogUrl := RelayConfig.EventLogUrl
	if len(EventLogUrl) == 0 {
		return
	}

	// 로그 기록
	evt := isc.EventLog{}
	evt.SetEventGatewayType("집중게이트웨이")
	evt.SetEventType("데이터 브로커")
	evt.SetEventGatewayId(RelayConfig.SourceId)
	t := time.Now()
	evt.SetEventTime(fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()))

	evt.SetEventStatus("점검")
	evt.SetEventMessage(err.Error())
	evt.SetEventContent(err.Error())
	isc.ReportLog(EventLogUrl, evt)

}

func main() {

	configPath := flag.String("c", "config.yaml", "configuration path(default:config.yaml)")
	srcId := flag.String("s", "center-gw-01", "gateway name")

	flag.Parse()

	// Confing
	LoadRelayConfiguration(*configPath)

	LocalTlsServerPort := RelayConfig.LocalTlsServerPort
	RemoteTlsServerIp := RelayConfig.RemoteTlsServerIp
	RemoteTlsServerPort := RelayConfig.RemoteTlsServerPort
	TlsCert := RelayConfig.TlsCert
	TlsKey := RelayConfig.TlsKey
	RelayConfig.SourceId = *srcId

	if &LocalTlsServerPort == nil || LocalTlsServerPort == 0 {
		fmt.Println("LocalTlsServerPort Required...")
		return
	}

	if &RemoteTlsServerIp == nil || len(RemoteTlsServerIp) == 0 {
		fmt.Println("RemoteTlsServerIp Required...")
		return
	}

	if &RemoteTlsServerPort == nil || RemoteTlsServerPort == 0 {
		fmt.Println("RemoteTlsServerPort Required...")
		return
	}

	if &RemoteTlsServerPort == nil || RemoteTlsServerPort == 0 {
		fmt.Println("RemoteTlsServerPort Required...")
		return
	}

	if &RemoteTlsServerPort == nil || RemoteTlsServerPort == 0 {
		fmt.Println("RemoteTlsServerPort Required...")
		return
	}

	isc.InitTLS2TLS(RemoteTlsServerIp, RemoteTlsServerPort, LocalTlsServerPort, TlsCert, TlsKey, callbck_handle_error)
}
