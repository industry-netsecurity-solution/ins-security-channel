package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
	"strconv"
)

var ClientConfig ins.ServiceConfigurations

var GConfigPath *string = nil
var GRemoteServerIp *string = nil
var GRemoteServerPort *int64 = nil

var Log *log.Logger = nil

/*
 *	parse command line
 */
func ParseOptions() int {
	GConfigPath = flag.String("c", "", "configuration path")
	GRemoteServerIp = flag.String("h", "", "remote address")
	GRemoteServerPort = flag.Int64("p", 0, "remote port")

	flag.Parse()

	return 0
}

func PrintConfig() {
	// Reading config ClientConfig
	for _, item := range  ClientConfig.ToString() {
		Log.Println(item)
	}
}

/*
 * load configuration
 */
func LoadConfiguration() int {
	if GConfigPath != nil && 0 < len(*GConfigPath){
		// Set the file name of the configurations file
		//viper.AddConfigPath(configFile)
		viper.SetConfigFile(*GConfigPath)
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			Log.Printf("Error reading config file, %s", err.Error())
		}

		err := viper.Unmarshal(&ClientConfig)
		if err != nil {
			Log.Printf("Unable to decode into struct, %v", err)
			return -1
		}
	}

	if GRemoteServerIp != nil && 0 < len(*GRemoteServerIp) {
		ClientConfig.Address = *GRemoteServerIp
	}

	if GRemoteServerPort != nil && 0 != *GRemoteServerPort {
		ClientConfig.Port = *GRemoteServerPort
	}

	if len(ClientConfig.Address) == 0 {
		Log.Println("Address Required...")
		return -1
	}

	if ClientConfig.Port == 0 {
		Log.Println("Port Required...")
		return -1
	}

	// Reading config file
	PrintConfig()

	return 0
}

func GetServerURL() *string {
	// TLS 서버 접속 정보
	if len(ClientConfig.Address) == 0 {
		Log.Println("Address Required...")
		return nil
	}

	if ClientConfig.Port == 0 {
		Log.Println("Port Required...")
		return nil
	}

	tlsServerURL := ClientConfig.Address + ":" + strconv.FormatInt(ClientConfig.Port, 10)

	return &tlsServerURL
}

func GetTCPAddr() (*net.TCPAddr, error) {
	// TLS 서버 주소
	serverURL := GetServerURL()
	if serverURL == nil {
		return nil, nil
	}

	rAddr, err := net.ResolveTCPAddr("tcp", *serverURL)
	if err != nil {
		Log.Println(err)
		return nil, err
	}

	return rAddr, nil
}

// 텔레필드 Tag
func DefaultTag() []byte {
	mesgType := []byte{0xff, 0x00}
	return mesgType
}

/**
 * A-brain 작업자 식별 메시지 전송
 */
func SendMessage(data []byte, conn net.Conn) (int, error) {

	var packet bytes.Buffer

	packet.Write(data)

	n, err := conn.Write(packet.Bytes())
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return n, nil
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

	serverUrl := ClientConfig.Address + ":" + strconv.FormatInt(ClientConfig.Port, 10)
	addr, err := net.ResolveTCPAddr("tcp", serverUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	var conn net.Conn = nil
	if ClientConfig.EnableTls {
		var config *tls.Config = nil
		config = &tls.Config{
			InsecureSkipVerify: true,
		}
		// TCP/TLS 연결
		conn, err = tls.Dial("tcp", addr.String(), config)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		conn, err = net.DialTCP("tcp", nil, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	defer conn.Close()

	message := "telefield message sample data"
	data := []byte(message)

	mesgType := []byte{0xFF, 0x00}
	packet := ins.EncTagLnV(binary.LittleEndian, mesgType, 32, data)

	n, err := SendMessage(packet, conn)
	if err != nil {
		panic(err)
	}

	Log.Printf("%d byte 메시지 파일 전송 성공", n)
}
