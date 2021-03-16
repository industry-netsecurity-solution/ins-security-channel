package main

import (
	"flag"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/config"
	"github.com/industry-netsecurity-solution/ins-security-channel/tls"
	"github.com/spf13/viper"
	"os"
)

var RelayConfig config.RelayConfigurations

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

func main() {

	configPath := flag.String("c", "config.yaml", "configuration path(default:config.yaml)")
	flag.Parse()

	// Confing
	LoadRelayConfiguration(*configPath)

	LocalTlsServerPort := RelayConfig.LocalTlsServerPort
	RemoteTlsServerIp := RelayConfig.RemoteTlsServerIp
	RemoteTlsServerPort := RelayConfig.RemoteTlsServerPort
	TlsCert := RelayConfig.TlsCert
	TlsKey := RelayConfig.TlsKey

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

	tls.InitTLS2TLS(RemoteTlsServerIp, RemoteTlsServerPort, LocalTlsServerPort, TlsCert, TlsKey)
}
