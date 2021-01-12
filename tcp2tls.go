package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
)

func initTCP2TLS(tlsServerIp string, tlsport int, tcpport int) {
	localTCP := "127.0.0.1:" + strconv.Itoa(tcpport)
	addr, err := net.ResolveTCPAddr("tcp", localTCP)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("Wait Accept...")
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Accepted:", conn.RemoteAddr())

		go func(connection net.Conn, tlsServerIp string, tlsport int) {
			go proxyConnectionTCP2TLS(connection, tlsServerIp, tlsport)
			fmt.Println("Send Server : ", tlsServerIp)
		}(conn, tlsServerIp, tlsport)
	}
}

func proxyConnectionTCP2TLS(conn net.Conn, tlsServerIp string, tlsport int) {
	defer conn.Close()

	tlsServerURL := tlsServerIp + ":" + strconv.Itoa(tlsport)
	rAddr, err := net.ResolveTCPAddr("tcp", tlsServerURL)
	if err != nil {
		fmt.Println(err)
	}

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	rConn, err := tls.Dial("tcp", rAddr.String(), conf)
	if err != nil {
		fmt.Println(err)
	}

	relay(conn, rConn)
	defer rConn.Close()

}
