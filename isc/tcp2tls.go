package isc

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
)

type ConnectionData struct {
	buffer1 bytes.Buffer
	buffer2 bytes.Buffer
}

/**
 * TCP connection을 TLS connection으로 변경하여 전달한다.
 */
func proxyConnection2TLS(conn net.Conn, relayConfig RelayConfigurations, callback_event func([]byte, []byte, *error)) {
	defer conn.Close()

	// TLS 서버 주소
	tlsServerURL := relayConfig.RemoteServerIp + ":" + strconv.Itoa(relayConfig.RemoteServerPort)
	rAddr, err := net.ResolveTCPAddr("tcp", tlsServerURL)
	if err != nil {
		if callback_event != nil {
			callback_event(nil, nil, &err)
		}
		fmt.Println(err)
		return
	}

	var rConn net.Conn = nil
	if relayConfig.EnableRemoteTls {
		var config *tls.Config = nil
		config = &tls.Config{
			InsecureSkipVerify: true,
		}
		// TCP/TLS 연결
		rConn, err = tls.Dial("tcp", rAddr.String(), config)
		if err != nil {
			if callback_event != nil {
				callback_event(nil, nil, &err)
			}
			fmt.Println(err)
			return
		}
	} else {

		rConn, err = net.DialTCP("tcp", nil, rAddr)
		if err != nil {
			if callback_event != nil {
				callback_event(nil, nil, &err)
			}
			fmt.Println(err)
			return
		}

	}

	// tcp conn <--> tls conn 간 상호 전달
	if Relay(conn, rConn, callback_event) == 0 {
		if callback_event != nil {
			callback_event(nil, nil, nil)
		}
	}

	defer rConn.Close()
}


/**
 * TCP connection을 TLS connection으로 변경하여 전달한다.
 */
func proxyConnectionEngageBroker(conn net.Conn,
	relayConfig RelayConfigurations,
	filter_event func(*ConnectionData, []byte, []byte, *error)([]byte, int),
	userdata *ConnectionData) {
	defer conn.Close()

	if filter_event == nil {
		return
	}

	// TLS 서버 주소
	tlsServerURL := relayConfig.RemoteServerIp + ":" + strconv.Itoa(relayConfig.RemoteServerPort)
	rAddr, err := net.ResolveTCPAddr("tcp", tlsServerURL)
	if err != nil {
		filter_event(userdata, nil, nil, &err)

		fmt.Println(err)
		return
	}

	var rConn net.Conn = nil
	if relayConfig.EnableRemoteTls {
		var config *tls.Config = nil
		config = &tls.Config{
			InsecureSkipVerify: true,
		}
		// TCP/TLS 연결
		rConn, err = tls.Dial("tcp", rAddr.String(), config)
		if err != nil {
			filter_event(userdata, nil, nil, &err)
			fmt.Println(err)
			return
		}
	} else {

		rConn, err = net.DialTCP("tcp", nil, rAddr)
		if err != nil {
			filter_event(userdata, nil, nil, &err)
			fmt.Println(err)
			return
		}

	}

	// tcp conn <--> tls conn 간 상호 전달
	if RelayBroker(conn, rConn, filter_event, userdata) == 0 {
		filter_event(userdata, nil, nil, nil)
	}

	defer rConn.Close()
}

/**
 * 로컬 수신 대기
 */
func InitTCP2TLS(relayConfig RelayConfigurations, callback_event func([]byte, []byte, *error)) {
	localurl := "0.0.0.0:" + strconv.Itoa(relayConfig.LocalServerPort)
	addr, err := net.ResolveTCPAddr("tcp", localurl)
	if err != nil {
		if callback_event != nil {
			callback_event(nil, nil, &err)
		}
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		if callback_event != nil {
			callback_event(nil, nil, &err)
		}
		panic(err)
	}

	for {
		fmt.Println("Wait Accept...")
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Accepted:", conn.RemoteAddr())

		fmt.Println("Send Server : ", relayConfig.RemoteServerIp)
		go proxyConnection2TLS(conn, relayConfig, callback_event)
	}
}

/**
 * 로컬 수신 대기
 */
func InitTLS2TLS(relayConfig RelayConfigurations, callback_event func([]byte, []byte, *error)) {
	localurl := "0.0.0.0:" + strconv.Itoa(relayConfig.LocalServerPort)
	//cer, err := tls.LoadX509KeyPair("server.crt", "server.key")

	var listener net.Listener = nil
	if relayConfig.EnableLocalTls {
		var config *tls.Config = nil
		cer, err := tls.LoadX509KeyPair(relayConfig.TlsCert, relayConfig.TlsKey)
		if err != nil {
			if callback_event != nil {
				callback_event(nil, nil, &err)
			}
			panic(err)
		}

		config = &tls.Config{Certificates: []tls.Certificate{cer}}
		listener, err = tls.Listen("tcp", localurl, config)
		if err != nil {
			if callback_event != nil {
				callback_event(nil, nil, &err)
			}
			panic(err)
		}
	} else {
		addr, err := net.ResolveTCPAddr("tcp", localurl)
		listener, err = net.ListenTCP("tcp", addr)
		if err != nil {
			if callback_event != nil {
				callback_event(nil, nil, &err)
			}
			panic(err)
		}
	}

	defer listener.Close()

	for {
		fmt.Println("Wait Accept...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Accepted:", conn.RemoteAddr())

		go proxyConnection2TLS(conn, relayConfig, callback_event)
	}
}


/**
 * 로컬 수신 대기
 */
func InitEngageTLS2TLS(relayConfig RelayConfigurations,
	filter_event func( *ConnectionData, []byte, []byte, *error)([]byte, int)) {
	localurl := "0.0.0.0:" + strconv.Itoa(relayConfig.LocalServerPort)
	//cer, err := tls.LoadX509KeyPair("server.crt", "server.key")

	var listener net.Listener = nil
	if relayConfig.EnableLocalTls {
		var config *tls.Config = nil
		cer, err := tls.LoadX509KeyPair(relayConfig.TlsCert, relayConfig.TlsKey)
		if err != nil {
			if filter_event != nil {
				filter_event(nil, nil, nil, &err)
			}
			panic(err)
		}

		config = &tls.Config{Certificates: []tls.Certificate{cer}}
		listener, err = tls.Listen("tcp", localurl, config)
		if err != nil {
			if filter_event != nil {
				filter_event(nil, nil, nil, &err)
			}
			panic(err)
		}
	} else {
		addr, err := net.ResolveTCPAddr("tcp", localurl)
		listener, err = net.ListenTCP("tcp", addr)
		if err != nil {
			if filter_event != nil {
				filter_event(nil, nil, nil, &err)
			}
			panic(err)
		}
	}

	defer listener.Close()

	for {
		fmt.Println("Wait Accept...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Accepted:", conn.RemoteAddr())

		userdata := new(ConnectionData)

		go proxyConnectionEngageBroker(conn, relayConfig, filter_event, userdata)
	}
}
