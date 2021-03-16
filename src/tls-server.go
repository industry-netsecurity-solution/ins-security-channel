package main

import (
	"bufio"
	"encoding/binary"
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
		fmt.Printf("Error reading config file, %s\n", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Printf("%s", dir)

	err = viper.Unmarshal(&ServerConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
	}

	// Reading config file
	fmt.Println("Reading config file")
	fmt.Println("LocalTlsServerPort is\t", ServerConfig.LocalTlsServerPort)
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

	file, err := os.Create(fileName)
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
