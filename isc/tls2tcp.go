package isc

import (
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
)

func proxyConnectionTLS2TCP(conn net.Conn, tcpport int, callback_event func([]byte, []byte, *error)) {
	localTCP := "127.0.0.1:" + strconv.Itoa(tcpport)
	defer conn.Close()

	rAddr, err := net.ResolveTCPAddr("tcp", localTCP)
	if err != nil {
		if callback_event != nil {
			callback_event(nil, nil, &err)
		}
		fmt.Println(err)
		return
	}

	rConn, err := net.DialTCP("tcp", nil, rAddr)
	if err != nil {
		if callback_event != nil {
			callback_event(nil, nil, &err)
		}
		fmt.Println(err)
		return
	}

	defer rConn.Close()

	Relay(conn, rConn, callback_event)

	fmt.Println("Handle Connection end...")
}

func InitTLS2TCP(tlsport int, tcpport int, callback_event func([]byte, []byte, *error)) {
	localTLS := strconv.Itoa(tlsport)
	cer, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		if callback_event != nil {
			callback_event(nil, nil, &err)
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
			if callback_event != nil {
				callback_event(nil, nil, &err)
			}
			fmt.Println(err)
			continue
		}

		go proxyConnectionTLS2TCP(conn, tcpport, callback_event)
	}
}

func ReadyServer(serverConfig *ServerConfigurations, callback func(net.Conn) error) int {

	if serverConfig == nil {
		return -1
	}

	if callback == nil {
		return -1
	}


	var listener net.Listener = nil

	localurl := "0.0.0.0:" + strconv.Itoa(serverConfig.LocalServerPort)
	if serverConfig.EnableLocalTls {
		var config *tls.Config = nil
		//cer, err := tls.LoadX509KeyPair("server.pem", "server.key")
		cer, err := tls.LoadX509KeyPair(serverConfig.TlsCert, serverConfig.TlsKey)
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

		go callback(conn)
	}

	return 0
}
