package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
)

var Config ins.RelayConfigurations

var GConfigPath *string = nil
var GSourceId *string = nil
var GDiag *int64 = nil

var GServiceAddress *string = nil
var GServicePort *int64 = nil
var GRemoteAddress *string = nil
var GRemotePort *int64 = nil

var Log *log.Logger = nil

/*
 *	parse command line
 */
func ParseFlagOptions() int {
	GConfigPath = flag.String("c", "", "configuration path")

	GSourceId = flag.String("source", "", "The mesage source id")
	GDiag = flag.Int64("diag", 0, "The diag interval")

	GServiceAddress = flag.String("service.address", "0.0.0.0", "The address of service. default: 0.0.0.0")
	GServicePort = flag.Int64("service.port", 9980, "The port of service. default: 9980")
	GRemoteAddress = flag.String("remote.address", "127.0.0.1", "The address to relay. default: 127.0.0.1")
	GRemotePort = flag.Int64("remote.port", 9990, "The port to relay. default: 9990")

	flag.Parse()

	return 0
}

/*
 * load configuration
 */
func LoadConfiguration() int {
	if GConfigPath != nil && 0 < len(*GConfigPath){
		viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

		// Set the file name of the configurations file
		//viper.AddConfigPath(configFile)
		viper.SetConfigFile(*GConfigPath)
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			Log.Printf("Error reading config file, %s", err.Error())
		}

		if viper.Get("DiagInterval") == nil {
			viper.SetDefault("DiagInterval", 60)
		}

		err := viper.Unmarshal(&Config)
		if err != nil {
			Log.Printf("Unable to decode into struct, %v", err)
			return -1
		}
	}

	return 0
}

/*
 * load configuration
 */
func UpdateConfiguration() int {

	// Service
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

	// Remote
	if ins.IsFlagPassed("remote.address") {
		Config.Remote.Address = *GRemoteAddress
	} else {
		if len(Config.Remote.Address) == 0 {
			Config.Remote.Address = *GRemoteAddress
		}
	}

	if ins.IsFlagPassed("remote.port") {
		Config.Remote.Port = *GRemotePort
	} else {
		if Config.Remote.Port == 0 {
			Config.Remote.Port = *GRemotePort
		}
	}

	//
	if ins.IsFlagPassed("source") {
		Config.SourceId = *GSourceId
	} else {
		if len(Config.SourceId) == 0 {
			Config.SourceId = *GSourceId
		}
	}

	if ins.IsFlagPassed("diag") {
		Config.DiagInterval = *GDiag
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

	return 0
}

/**
 * 집중 GW와 연동 메시지로 변경한다.
 */
func make2CenterGW(data []byte) bytes.Buffer {

	buffer := bytes.Buffer{}
	buffer.Write(data)

	return buffer
}

/*
func cb_session(conn net.Conn, data []byte, err error, ud interface{}) ([]byte, error) {
	buffer := makeWrapPacket0xFF00(data)

	fmt.Printf("Relay: %d -> %d byte\n", len(data), buffer.Len())
	return buffer.Bytes(), err
}*/

func cb_session(conn net.Conn, ud interface{}) ([]byte, error) {
	data, err := ins.RecvTLV(conn, binary.LittleEndian)
	if err != nil {
		return nil, err
	}
	buffer := make2CenterGW(data)

	fmt.Printf("Reform: %d -> %d byte\n", len(data), buffer.Len())
	return buffer.Bytes(), err
}

func callback(conn1 net.Conn, ud interface{}) error {

	defer conn1.Close()

	conn2, err := ins.Dial(&Config.Remote)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer conn2.Close()

	ins.Relay(conn1 /* net.Conn */,
		nil /* ud1 interface{} */,
		cb_session /* func(net.Conn, interface{}) ([]byte, error) */,
		conn2 /* net.Conn */,
		nil /* ud2 interface{} */,
		cb_session /* func(net.Conn, interface{}) ([]byte, error) */)

	return nil
}

func main() {

	//MQTT.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	//MQTT.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	//MQTT.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	//MQTT.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	Log = log.New(os.Stdout, "", log.LstdFlags)

	// 프로그램 인자 및 Configuration 확인
	if ParseOptions() == -1 {
		return
	}

	ins.ReadyServer(&Config.Service, nil, callback)
}
