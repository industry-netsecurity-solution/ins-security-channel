package isc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"net"
	"time"
)

func CopyFromTCP(conn net.Conn, callback_event func([]byte, []byte, *error)) chan []byte {
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
				if err == io.EOF {
					c <- nil
					break
				}

				if callback_event != nil {
					callback_event(nil, nil, &err)
				}
				c <- nil
				break
			}
		}
	}()

	return c
}

func FilterFromTCP(conn net.Conn,
	callback_event func(*ConnectionData, []byte, []byte, *error)([]byte, int),
	userdata *ConnectionData) chan []byte {
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
				if err == io.EOF {
					c <- nil
					break
				}

				if callback_event != nil {
					callback_event(userdata, nil, nil, &err)
				}
				c <- nil
				break
			}
		}
	}()

	return c
}

/**
 *  있는 그대로 읽고 전달한다.
 */
func Relay(conn1 net.Conn, conn2 net.Conn, callback_event func([]byte, []byte, *error)) int {
	var len1 int = -1
	var len2 int = -1

	chan1 := CopyFromTCP(conn1, callback_event)
	chan2 := CopyFromTCP(conn2, callback_event)

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

/**
 * 데이터를 받아서, 처리하고 결과를 전달한다.
 */
func RelayBroker(conn1 net.Conn, conn2 net.Conn,
	broker_event func(*ConnectionData, []byte, []byte, *error) ([]byte, int),
	userdata *ConnectionData) int {

	if broker_event == nil {
		return -1
	}

	var len1 int = -1
	var len2 int = -1

	chan1 := FilterFromTCP(conn1, broker_event, userdata)
	chan2 := FilterFromTCP(conn2, broker_event, userdata)

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

				// conn1에서 수신한 데이터
				data, nlen := broker_event(userdata, b1, nil, nil)
				if 0 < nlen {
					_, err := conn2.Write(data)
					if err != nil {
						broker_event(userdata, nil, nil, &err)
						return -1
					}
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

				// conn2에서 수신한 데이터
				data, nlen := broker_event(userdata, nil, b2, nil)
				if 0 < nlen {
					_, err := conn1.Write(data)
					if err != nil {
						broker_event(userdata, nil, nil, &err)
						return -1
					}
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
	client.SetCloseConnection(true)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(url)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(resp)

	// resp.RawBody().Close()

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

func EncTagLn(order binary.ByteOrder, tag []byte, size uint, length uint32) []byte {
	var buffer = new(bytes.Buffer)
	binary.Write(buffer, order, tag)

	if size == 16 {
		binary.Write(buffer, order, uint16(length))
	} else if size == 32 {
		binary.Write(buffer, order, uint32(length))
	} else if size == 64 {
		binary.Write(buffer, order, uint64(length))
	} else {
		return nil
	}

	return buffer.Bytes()
}

func EncTagLnV(order binary.ByteOrder, tag []byte, size uint, data []byte) []byte {
	var buffer = new(bytes.Buffer)
	binary.Write(buffer, order, tag)

	if data == nil {
		if size == 16 {
			binary.Write(buffer, order, uint16(0))
		} else if size == 32 {
			binary.Write(buffer, order, uint32(0))
		} else if size == 64 {
			binary.Write(buffer, order, uint64(0))
		} else {
			return nil
		}

		return buffer.Bytes()
	}

	if size == 16 {
		binary.Write(buffer, order, uint16(len(data)))
	} else if size == 32 {
		binary.Write(buffer, order, uint32(len(data)))
	} else if size == 64 {
		binary.Write(buffer, order, uint64(len(data)))
	} else {
		return nil
	}

	buffer.Write(data)

	return buffer.Bytes()
}

func EncTagLnString(order binary.ByteOrder, tag []byte, size uint, data string) []byte {
	return EncTagLnV(order, tag, size, []byte(data))
}

func EncTagLnUInt32(order binary.ByteOrder, tag []byte, size uint, data uint32) []byte {
	buf4 := make([]byte, 4)
	order.PutUint32(buf4, data)
	return EncTagLnV(order, tag, size, buf4)
}

/*
 * LittleEndian
 */
func EncLETL(tag uint16, length uint32) []byte {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	var buffer bytes.Buffer

	binary.LittleEndian.PutUint16(buf2, tag)
	buffer.Write(buf2)

	binary.LittleEndian.PutUint32(buf4, length)
	buffer.Write(buf4)

	return buffer.Bytes()
}

func EncLETLnV(tag uint16, size uint, data []byte) []byte {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)
	buf8 := make([]byte, 8)

	var buffer bytes.Buffer

	binary.LittleEndian.PutUint16(buf2, tag)
	buffer.Write(buf2)

	if data == nil {
		if size == 16 {
			binary.LittleEndian.PutUint16(buf2, uint16(0))
			buffer.Write(buf2)
		} else if size == 32 {
			binary.LittleEndian.PutUint32(buf4, uint32(0))
			buffer.Write(buf4)
		} else if size == 64 {
			binary.LittleEndian.PutUint64(buf8, uint64(0))
			buffer.Write(buf8)
		} else {
			return nil
		}

		return buffer.Bytes()
	}

	if size == 16 {
		binary.LittleEndian.PutUint16(buf2, uint16(len(data)))
		buffer.Write(buf2)
	} else if size == 32 {
		binary.LittleEndian.PutUint32(buf4, uint32(len(data)))
		buffer.Write(buf4)
	} else if size == 64 {
		binary.LittleEndian.PutUint64(buf8, uint64(len(data)))
		buffer.Write(buf8)
	} else {
		return nil
	}

	buffer.Write(data)

	return buffer.Bytes()
}

func EncLETLV(tag uint16, data []byte) []byte {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	var buffer bytes.Buffer

	binary.LittleEndian.PutUint16(buf2, tag)
	buffer.Write(buf2)

	if data == nil {
		binary.LittleEndian.PutUint32(buf4, uint32(0))
		buffer.Write(buf4)
		return buffer.Bytes()
	}

	binary.LittleEndian.PutUint32(buf4, uint32(len(data)))
	buffer.Write(buf4)

	buffer.Write(data)

	return buffer.Bytes()
}

func EncLEString(tag uint16, data string) []byte {
	return EncLETLV(tag, []byte(data))
}

func EncLEUint16(tag uint16, data uint16) []byte {
	buf2 := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf2, data)

	return EncLETLV(tag, buf2)
}

func EncLEUint32(tag uint16, data uint32) []byte {
	buf4 := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf4, data)

	return EncLETLV(tag, buf4)
}

/*
 * BigEndian
 */
func EncBETL(tag uint16, length uint32) []byte {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	var buffer bytes.Buffer

	binary.BigEndian.PutUint16(buf2, tag)
	buffer.Write(buf2)

	binary.BigEndian.PutUint32(buf4, length)
	buffer.Write(buf4)

	return buffer.Bytes()
}

func EncBETLV(tag uint16, data []byte) []byte {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	var buffer bytes.Buffer

	binary.BigEndian.PutUint16(buf2, tag)
	buffer.Write(buf2)

	if data == nil {
		binary.BigEndian.PutUint32(buf4, uint32(0))
		buffer.Write(buf4)
		return buffer.Bytes()
	}

	binary.BigEndian.PutUint32(buf4, uint32(len(data)))
	buffer.Write(buf4)

	buffer.Write(data)

	return buffer.Bytes()
}

func EncBEString(tag uint16, data string) []byte {
	return EncBETLV(tag, []byte(data))
}

func EncBEUint16(tag uint16, data uint16) []byte {
	buf2 := make([]byte, 2)
	binary.BigEndian.PutUint16(buf2, data)

	return EncBETLV(tag, buf2)
}

func EncBEUint32(tag uint16, data uint32) []byte {
	buf4 := make([]byte, 4)
	binary.BigEndian.PutUint32(buf4, data)

	return EncBETLV(tag, buf4)
}
