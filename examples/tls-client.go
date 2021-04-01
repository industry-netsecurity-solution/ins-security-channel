package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"flag"
	"fmt"
	isc "github.com/industry-netsecurity-solution/ins-security-channel/isc"
	"github.com/spf13/viper"
	"io"
	"log"
	"net"
	"os"
	"path"
	"strconv"
)

var ClientConfig isc.ClientConfigurations

var GConfigPath *string = nil
var GSrcId *string = nil
var GOptname *string = nil
var GFileType *uint = nil
var GArgs []string = nil

var Log *log.Logger = nil

/*
 *	parse command line
 */
func ParseOptions() int {
	GConfigPath = flag.String("c", "config.yaml", "configuration path(default:config.yaml)")
	GFileType = flag.Uint("f", 0, "file type (0: normal, 1: front video, 2: rear video, 3: front collision, 4: rear collision, 5: front approach, 6: rear approach)")
	GSrcId = flag.String("s", "bb-01", "gateway name")
	GOptname = flag.String("n", "", "gateway name")

	flag.Parse()

	if GFileType == nil {
		Log.Println("FileType Required...")
		return -1
	}

	GArgs = flag.Args()
	if GArgs == nil || len(GArgs) == 0 {
		Log.Println("Filename Required...")
		return -1
	}

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
	if GConfigPath == nil {
		return -1
	}

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

	if len(ClientConfig.RemoteServerIp) == 0 {
		Log.Println("RemoteTlsServerIp Required...")
		return -1
	}

	if ClientConfig.RemoteServerPort == 0 {
		Log.Println("RemoteTlsServerPort Required...")
		return -1
	}

	// Reading config file
	PrintConfig()

	return 0
}

func GetServerURL() *string {
	// TLS 서버 접속 정보
	if len(ClientConfig.RemoteServerIp) == 0 {
		Log.Println("RemoteTlsServerIp Required...")
		return nil
	}

	if ClientConfig.RemoteServerPort == 0 {
		Log.Println("RemoteTlsServerPort Required...")
		return nil
	}

	tlsServerURL := ClientConfig.RemoteServerIp + ":" + strconv.Itoa(ClientConfig.RemoteServerPort)

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

// 유미테크 Tag
func DefaultTag() []byte {
	ymtechType := []byte{0xef, 0xf0}
	return ymtechType
}

/**
 * file 전송
 */
func SendFileOverTls(file *os.File, params map[int][]byte, conn *tls.Conn) error {

	var buffer bytes.Buffer
	var payloadBuffer bytes.Buffer

	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	// EFF0  {Length}, {payload}
	// payload: [0, {source}], [1, {file type}], [2, {file name}], [3, {file size}], [4, {file contents}],

	// file size
	fileSize := binary.LittleEndian.Uint32(params[3])

	// [0, {source}], [1, {file type}], [2, {file name}], [3, {file size}]
	encodedParams := isc.EncodeMap(params)
	payloadBuffer.Write(encodedParams)

	// [4, {file contents}],
	// TAG: file contents
	binary.LittleEndian.PutUint16(buf2, uint16(4))
	payloadBuffer.Write(buf2)

	// LENGTH: file size
	binary.LittleEndian.PutUint32(buf4, uint32(fileSize))
	payloadBuffer.Write(buf4)

	// file 내용을 제외한 데이터
	data := payloadBuffer.Bytes()

	payloadLength := uint32(len(data)) + fileSize

	binary.LittleEndian.PutUint32(buf4, uint32(payloadLength))

	// 전체 패킷
	tag := DefaultTag()
	buffer.Write(tag)
	buffer.Write(buf4)
	buffer.Write(data)

	_, err := conn.Write(buffer.Bytes())
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 실제 파일 내용
	for {
		buf := make([]byte, 4096)
		rlen, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return err
		}

		res := make([]byte, rlen)
		copy(res, buf[:rlen])

		_, err = conn.Write(res)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func sendfile() {

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

	rAddr, _ := GetTCPAddr()
	if rAddr == nil {
		return
	}

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	filename := GArgs[0]

	var sendname string
	if GOptname != nil && 0 < len(*GOptname) {
		Log.Printf("send: %s -> %s\n", filename, *GOptname)
		sendname = *GOptname
	} else {
		Log.Printf("send: %s", filename)
		sendname = filename
	}

	stat, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}

	file, err := os.Open(filename)
	defer file.Close()

	// TLS 연결
	tlsConn, err := tls.Dial("tcp", rAddr.String(), conf)
	if err != nil {
		panic(err)
	}

	defer tlsConn.Close()

	params := make(map[int][]byte)

	// source
	params[0] = []byte(*GSrcId)

	// file type
	params[1] = make([]byte, 2)
	binary.LittleEndian.PutUint16(params[1], uint16(*GFileType))

	// file name
	params[2] = []byte(path.Base(sendname))

	// file size
	params[3] = make([]byte, 4)
	binary.LittleEndian.PutUint32(params[3], uint32(stat.Size()))

	err = SendFileOverTls(file, params, tlsConn)
	if err != nil {
		panic(err)
	}

	Log.Printf("%s 파일 전송 성공\n", sendname)
}
