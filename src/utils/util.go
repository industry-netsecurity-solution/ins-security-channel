package utils

import (
	"crypto/tls"
	"fmt"
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

func Relay(conn1 net.Conn, conn2 net.Conn) {
	chan1 := CopyFromTCP(conn1)
	chan2 := CopyFromTCP(conn2)

	for {
		select {
		case b1 := <-chan1:
			if b1 == nil {
				return
			} else {
				conn2.Write(b1)
			}
		case b2 := <-chan2:
			if b2 == nil {
				return
			} else {
				conn1.Write(b2)
			}
		}
	}
}

/**
 * file 전송
 */
func SendFileOverTls(file *os.File, conn *tls.Conn) error {




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
