package ins

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

func NewTLSConfig() *tls.Config {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to
	// default openssl CA bundle.
	certpool := x509.NewCertPool()
	//pemCerts, err := ioutil.ReadFile("mosquitto-ca.crt")
	//if err == nil {
	//	certpool.AppendCertsFromPEM(pemCerts)
	//}

	// Import client certificate/key pair
	//cert, err := net.LoadX509KeyPair("samplecerts/client-crt.pem", "samplecerts/client-key.pem")
	//if err != nil {
	//	panic(err)
	//}

	// Just to print out the client certificate..
	//cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(cert.Leaf)

	// Create net.Config with desired net properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		//Certificates: []net.Certificate{cert},
	}
}

func RecvTLV(conn net.Conn, order binary.ByteOrder) ([]byte, error) {
	buf := bytes.Buffer{}
	tagArray := make([]byte, 2)
	lengthArray := make([]byte, 4)

	n, err := conn.Read(tagArray)
	if err != nil {
		return nil, err
	}
	if n != 2 {
		return nil, errors.New("The tag is not enough length.")
	}
	buf.Write(tagArray)

	n, err = conn.Read(lengthArray)
	if err != nil {
		return nil, err
	}
	if n != 4 {
		return nil, errors.New("The length is not enough length.")
	}
	buf.Write(lengthArray)

	length := order.Uint32(lengthArray)

	if length == 0 {
		return buf.Bytes(), nil
	}

	dataArray := make([]byte, 4096)
	dataLength := 0

	for {
		if int(length) <= dataLength {
			break
		}
		if int(length) - dataLength < len(dataArray) {
			r := int(length) - dataLength
			n, err = conn.Read(dataArray[:r])
		} else {
			n, err = conn.Read(dataArray)
		}

		if 0 < n {
			buf.Write(dataArray[:n])
			dataLength += n
		}

		if err != nil {
			// 접속 종료
			if err == io.EOF {
				break
			}

			// 기타 오류
			return nil, err
		}
	}

	if dataLength < int(length) {
		return nil, errors.New("The value is not enough length.")
	}

	return buf.Bytes(), nil
}

func ReadyServer(serviceConfig *ServiceConfigurations, ud interface{}, callback func(net.Conn, interface{}) error) int {

	if serviceConfig == nil {
		return -1
	}

	if callback == nil {
		return -1
	}

	var listener net.Listener = nil
	var localurl string

	if len(serviceConfig.Address) == 0 {
		localurl = fmt.Sprintf("0.0.0.0:%d", serviceConfig.Port)
	} else {
		localurl = fmt.Sprintf("%s:%d", serviceConfig.Address, serviceConfig.Port)
	}

	if serviceConfig.EnableTls {
		var config *tls.Config = nil
		//cer, err := tls.LoadX509KeyPair("server.pem", "server.key")
		cer, err := tls.LoadX509KeyPair(serviceConfig.TlsCert, serviceConfig.TlsKey)
		if err != nil {
			panic(err)
			return -1
		}

		config = &tls.Config{Certificates: []tls.Certificate{cer}}
		listener, err = tls.Listen("tcp", localurl, config)
		if err != nil {
			panic(err)
			return -1
		}
	} else {
		addr, err := net.ResolveTCPAddr("tcp", localurl)
		listener, err = net.ListenTCP("tcp", addr)
		if err != nil {
			panic(err)
			return -1
		}
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go callback(conn, ud)
	}

	return 0
}

/**
 * remote 서비스에 연결을 시도한다.
 * timeout 연결 대기 시간으로 <= 0 면, timeout을 지정하지 않는다.
 */
func Dial(remote *ServiceConfigurations, timeout time.Duration) (net.Conn, error) {
	url := fmt.Sprintf("%s:%d", remote.Address, remote.Port)
	addr, err := net.ResolveTCPAddr("tcp", url)
	if err != nil {
		return nil, err
	}

	var dialer *net.Dialer = nil
	if 0 < timeout {
		dialer := new(net.Dialer)
		dialer.Timeout = timeout
	}

	var conn net.Conn = nil
	if remote.EnableTls {
		var config *tls.Config = nil
		config = &tls.Config{
			InsecureSkipVerify: true,
		}

		// TCP/TLS 연결
		if dialer == nil {
			conn, err = tls.Dial("tcp", addr.String(), config)
		} else {
			conn, err = tls.DialWithDialer(dialer, "tcp", addr.String(), config)
		}
		if err != nil {
			return nil, err
		}
	} else {
		if dialer == nil {
			conn, err = net.DialTCP("tcp", nil, addr)
		} else {
			conn, err = dialer.Dial("tcp", url)
		}
		if err != nil {
			return nil, err
		}
	}

	return conn, nil
}

/**
 *
 */
func MakeRecvChannel(conn net.Conn, ud interface{}, callback func(net.Conn, interface{}) ([]byte, error)) (chan []byte, chan error) {

	c := make(chan []byte)
	e := make(chan error)

	go func() {
		b := make([]byte, 4096)
		for {
			if callback == nil {
				n, err := conn.Read(b)
				if n > 0 {
					res := make([]byte, n)
					copy(res, b[:n])
					c <- res
				}

				if err != nil {
					e <- err

					close(c)
					close(e)
					return
				}
			} else {
				data, err := callback(conn, ud)
				if data != nil {
					c <- data
				}

				if err != nil {
					e <- err

					close(c)
					close(e)
					return
				}
			}
		}
	}()

	return c, e
}

/**
 *  있는 그대로 읽고 전달한다.
 */
func Relay(conn1 net.Conn, ud1 interface{},
	callback1 func(net.Conn, interface{}) ([]byte, error),
	conn2 net.Conn, ud2 interface{},
	callback2 func(net.Conn, interface{}) ([]byte, error)) int {

	chan1, echan1 := MakeRecvChannel(conn1, ud1, callback1)
	chan2, echan2 := MakeRecvChannel(conn2, ud2, callback2)

	for {
		select {
		case b1 := <- chan1:
			if b1 == nil {
				return 0
			} else {
				// 수신한 데이터를 상대 소켓에 보낸다.
				_, err := conn2.Write(b1)
				if err != nil {
					return -1
				}
			}
			break
		case e1 := <- echan1:
			if e1 == io.EOF {
				return 0
			}
			return -1
		case b2 := <-chan2:
			if b2 == nil {
				return 0
			} else {
				// 수신한 데이터를 상대 소켓에 보낸다.
				_, err := conn1.Write(b2)
				if err != nil {
					return -1
				}
			}
			break
		case e2 := <- echan2:
			if e2 == io.EOF {
				return 0
			}
			return -1
		}
	}

	return 0
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