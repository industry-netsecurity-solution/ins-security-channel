package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/config"
	"github.com/industry-netsecurity-solution/ins-security-channel/utils"
	"github.com/spf13/viper"
	"net"
	"os"
	"strconv"
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

	flag.Parse()

	// 프로그램 인자 확인
	args := flag.Args()

	if args == nil || len(args) == 0 {
		fmt.Println("Filename Required...")
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

	for index, filename := range args {
		fmt.Println(index, filename)

		file, err := os.Open(filename)
		defer file.Close()

		// TLS 연결
		tlsConn, err := tls.Dial("tcp", rAddr.String(), conf)
		if err != nil {
			//fmt.Println(err)
			panic(err)
			break
		}

		err = utils.SendFileOverTls(file, tlsConn)
		if err != nil {
			panic(err)
			break
		}
	}
}
