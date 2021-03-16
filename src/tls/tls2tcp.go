package tls

import (
	"crypto/tls"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/utils"
	"net"
	"strconv"
)

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

	utils.Relay(conn, rConn)

	fmt.Println("Handle Connection end...")
}

func InitTLS2TCP(tlsport int, tcpport int) {
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

func ReadyTLSServer(tlsport int, tlsCert string, tlskey string, callback func(net.Conn) error ) {

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
