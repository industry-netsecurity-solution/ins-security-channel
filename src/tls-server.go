package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/config"
	"github.com/industry-netsecurity-solution/ins-security-channel/tls"
	"github.com/spf13/viper"
	"net"
	"os"
)

var ServerConfig config.ServerConfigurations

func LoadServerConfiguration(configFile string) {
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

	err = viper.Unmarshal(&ServerConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading config file
	fmt.Println("Reading config file")
	fmt.Println("LocalTlsServerPort is\t", ServerConfig.LocalTlsServerPort)
}

/**
 * 연결로 부터 데이터를 읽어 파일로 저장한다.
 */
func receiveData(conn net.Conn) error {

	file, err := os.Create("/tmp/dat2")
	if err != nil {
		fmt.Println(err)
		return err
	}

	writer := bufio.NewWriter(file)

	defer func() {
		writer.Flush()
		file.Close()
	} ()

	for {
		buf := make([]byte, 4096)
		rlen, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Printf("rlen = %d\n", rlen)

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

	configPath := flag.String("c", "config.yaml", "configuration path(default:config.yaml)")
	flag.Parse()

	LoadServerConfiguration(*configPath)

	LocalTlsServerPort := ServerConfig.LocalTlsServerPort
	TlsCert := ServerConfig.TlsCert
	TlsKey := ServerConfig.TlsKey

	if &LocalTlsServerPort == nil || LocalTlsServerPort == 0 {
		fmt.Println("LocalTlsServerPort Required...")
		return
	}

	tls.ReadyTLSServer(LocalTlsServerPort, TlsCert, TlsKey, receiveData)
}
