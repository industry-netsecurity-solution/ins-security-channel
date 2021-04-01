package isc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net"
	"time"
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

func Relay(conn1 net.Conn, conn2 net.Conn, callback_event func([]byte, []byte, *error)) int {
	var len1 int = -1
	var len2 int = -1

	chan1 := CopyFromTCP(conn1)
	chan2 := CopyFromTCP(conn2)

	for {
		select {
		case b1 := <-chan1:
			if b1 == nil {
				return 0
			} else {
				if len1 < 0 {
					len1 = len(b1)
				} else {
					len1 += len(b1)
				}

				// 수신한 데이터
				if callback_event != nil {
					callback_event(b1, nil, nil)
				}

				_, err := conn2.Write(b1)
				if err != nil {
					if callback_event != nil {
						callback_event(nil, nil, &err)
					}

					return -1
				}
			}
		case b2 := <-chan2:
			if b2 == nil {
				return 0
			} else {
				if len2 < 0 {
					len2 = len(b2)
				} else {
					len2 += len(b2)
				}

				if callback_event != nil {
					callback_event(nil, b2, nil)
				}

				_, err := conn1.Write(b2)
				if err != nil {
					if callback_event != nil {
						callback_event(nil, nil, &err)
					}

					return -1
				}
			}
		}
	}

	return 0
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

func ReportLog(url string, log EventLog) error {

	//data, _e := json.Marshal(log)
	data, _e := log.EventLog()
	if _e != nil {
		return _e
	}

	fmt.Println(string(data))

	TransferHttp(url, data)
	return nil

}

func TimeYYmmddHHMMSS(t *time.Time) string {
	if t == nil {
		tm := time.Now()
		t = &tm;
	}

	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
