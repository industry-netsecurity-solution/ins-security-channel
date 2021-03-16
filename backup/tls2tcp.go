package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
)

func initTLS2TCP(tlsport int, tcpport int) {
	localTLS := strconv.Itoa(tlsport)
	cer, err := tls.LoadX509KeyPair("server.pem", "server.key")
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

		go proxyConnectionTLS2TCP(conn, tcpport)
	}
}

func proxyConnectionTLS2TCP(conn net.Conn, tcpport int) {
	localTCP := "127.0.0.1:" + strconv.Itoa(tcpport)
	defer conn.Close()

	rAddr, err := net.ResolveTCPAddr("tcp", localTCP)
	if err != nil {
		fmt.Println(err)
	}

	rConn, err := net.DialTCP("tcp", nil, rAddr)
	if err != nil {
		fmt.Println(err)
	}

	defer rConn.Close()

	relay(conn, rConn)

	fmt.Println("Handle Connection end...")
}
