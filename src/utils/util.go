package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/industry-netsecurity-solution/ins-security-channel/common"
	"io"
	"net"
	"os"
)

func CopyFromTCP(conn net.Conn) chan []byte {
	c := make(chan []byte)

	go func() {
		b := make([]byte, 4096)
		for {
			n, err := conn.Read(b)
			if n > 0 {
				res := make([]byte, n)
				copy(res, b[:n])
				c <- res
			}

			if err != nil {
				c <- nil
				break
			}
		}
	}()

	return c
}

func Relay(conn1 net.Conn, conn2 net.Conn, callback_error func(err error)) {
	chan1 := CopyFromTCP(conn1)
	chan2 := CopyFromTCP(conn2)

	for {
		select {
		case b1 := <-chan1:
			if b1 == nil {
				return
			} else {
				_, err := conn2.Write(b1)
				if err != nil {
					if callback_error != nil {
						callback_error(err)
					}
				}
			}
		case b2 := <-chan2:
			if b2 == nil {
				return
			} else {
				_, err := conn1.Write(b2)
				if err != nil {
					if callback_error != nil {
						callback_error(err)
					}
				}
			}
		}
	}
}

// 유미테크 Tag
func DefaultTag() []byte {
	ymtechType := []byte{0xef, 0xf0}
	return ymtechType
}

func EncodeMap(params map[int][]byte) []byte {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	var buffer bytes.Buffer

	for key := range params {

		// encode key
		binary.LittleEndian.PutUint16(buf2, uint16(key))
		buffer.Write(buf2)

		// encode length
		paramLen := len(params[key])
		binary.LittleEndian.PutUint32(buf4, uint32(paramLen))
		buffer.Write(buf4)

		// encode data
		buffer.Write(params[key])
	}

	return buffer.Bytes()
}

// HTTP를 통한 데이터 전송
func TransferHttp(url string, data []byte) error {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(url)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(resp)

	//resp.RawBody().Close()

	return nil
}

func ReportLog(log common.EventLog) error {

	//data, _e := json.Marshal(log)
	data, _e := log.EventLog()
	if _e != nil {
		return _e
	}

	fmt.Println(string(data))

	TransferHttp("http://10.10.1.15:9091/event/gateway", data)
	return nil

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
	encodedParams := EncodeMap(params)
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

	//
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
