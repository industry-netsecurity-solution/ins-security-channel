package main

import (
	"fmt"
	"github.com/spf13/viper"
)

var CONFIGURATION Configurations

func main() {
	configuration("config.yaml")

	RelayTcpServerPort := CONFIGURATION.RelayTcpServerPort
	RelayTlsServerPort := CONFIGURATION.RelayTlsServerPort
	RemoteTlsServerIp := CONFIGURATION.RemoteTlsServerIp
	RemoteTlsServerPort := CONFIGURATION.RemoteTlsServerPort
	LocalTcpServerPort := CONFIGURATION.LocalTcpServerPort

	if &RelayTcpServerPort == nil || RelayTlsServerPort == 0 {
		fmt.Println("RelayTcpServerPort Required...")
		return
	}

	if &RelayTlsServerPort == nil || RelayTlsServerPort == 0 {
		fmt.Println("RelayTlsServerPort Required...")
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

	if &LocalTcpServerPort == nil || LocalTcpServerPort == 0 {
		fmt.Println("LocalTcpServerPort Required...")
		return
	}

	initTCP2TLS(RemoteTlsServerIp, RemoteTlsServerPort, RelayTcpServerPort)
	initTLS2TCP(RelayTlsServerPort, LocalTcpServerPort)
}

func configuration(configFile string) {
	// Set the file name of the configurations file
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&CONFIGURATION)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading config file
	fmt.Println("Reading config file")
	fmt.Println("RelayTcpServerPort is\t", CONFIGURATION.RelayTcpServerPort)
	fmt.Println("RelayTlsServerPort is\t\t", CONFIGURATION.RelayTlsServerPort)
	fmt.Println("RemoteTlsServerIp is\t\t", CONFIGURATION.RemoteTlsServerIp)
	fmt.Println("RemoteTlsServerPort is\t\t", CONFIGURATION.RemoteTlsServerPort)
	fmt.Println("LocalTcpServerPort is\t\t", CONFIGURATION.LocalTcpServerPort)
}
