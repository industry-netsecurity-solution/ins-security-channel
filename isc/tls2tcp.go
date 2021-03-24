package isc

import (
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
)

func proxyConnectionTLS2TCP(conn net.Conn, tcpport int, callback_error func(err error)) {
	localTCP := "127.0.0.1:" + strconv.Itoa(tcpport)
	defer conn.Close()

	rAddr, err := net.ResolveTCPAddr("tcp", localTCP)
	if err != nil {
		if callback_error != nil {
			callback_error(err)
		}
		fmt.Println(err)
		return
	}

	rConn, err := net.DialTCP("tcp", nil, rAddr)
	if err != nil {
		if callback_error != nil {
			callback_error(err)
		}
		fmt.Println(err)
		return
	}

	defer rConn.Close()

	Relay(conn, rConn, callback_error)

	fmt.Println("Handle Connection end...")
}

func InitTLS2TCP(tlsport int, tcpport int, callback_error func(err error)) {
	localTLS := strconv.Itoa(tlsport)
	cer, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		if callback_error != nil {
			callback_error(err)
		}
		panic(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", localTLS, config)
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			if callback_error != nil {
				callback_error(err)
			}
			fmt.Println(err)
			continue
		}

		go proxyConnectionTLS2TCP(conn, tcpport, callback_error)
	}
}

func ReadyTLSServer(tlsport int, tlsCert string, tlskey string, callback func(net.Conn) error) {

	if callback == nil {
		return
	}

	localTLS := "0.0.0.0:" + strconv.Itoa(tlsport)
	//cer, err := tls.LoadX509KeyPair("server.pem", "server.key")
	cer, err := tls.LoadX509KeyPair(tlsCert, tlskey)
	if err != nil {
		panic(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", localTLS, config)
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go callback(conn)
	}
}
