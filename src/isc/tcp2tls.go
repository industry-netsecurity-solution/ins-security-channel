package isc

import (
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
)

/**
 * TCP connection을 TLS connection으로 변경하여 전달한다.
 */
func proxyConnection2TLS(conn net.Conn, tlsServerIp string, tlsport int, callback_error func(err error)) {
	defer conn.Close()

	// TLS 서버 주소
	tlsServerURL := tlsServerIp + ":" + strconv.Itoa(tlsport)
	rAddr, err := net.ResolveTCPAddr("tcp", tlsServerURL)
	if err != nil {
		if callback_error != nil {
			callback_error(err)
		}
		fmt.Println(err)
		return
	}

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	// TLS 연결
	rConn, err := tls.Dial("tcp", rAddr.String(), conf)
	if err != nil {
		if callback_error != nil {
			callback_error(err)
		}
		fmt.Println(err)
		return
	}

	// tcp conn <--> tls conn 간 상호 전달
	Relay(conn, rConn, callback_error)
	defer rConn.Close()

}

/**
 * 로컬 수신 대기
 */
func InitTCP2TLS(tlsServerIp string, tlsport int, tcpport int, callback_error func(err error)) {
	localTCP := "127.0.0.1:" + strconv.Itoa(tcpport)
	addr, err := net.ResolveTCPAddr("tcp", localTCP)
	if err != nil {
		if callback_error != nil {
			callback_error(err)
		}
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		if callback_error != nil {
			callback_error(err)
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

		fmt.Println("Send Server : ", tlsServerIp)
		go proxyConnection2TLS(conn, tlsServerIp, tlsport, callback_error)
	}
}

/**
 * 로컬 수신 대기
 */
func InitTLS2TLS(remoteTlsServerIp string, remoteTlsPort int, localTlsPort int, tlsCert string, tlskey string, callback_error func(err error)) {
	localTLS := "0.0.0.0:" + strconv.Itoa(localTlsPort)
	//cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	cer, err := tls.LoadX509KeyPair(tlsCert, tlskey)
	if err != nil {
		if callback_error != nil {
			callback_error(err)
		}
		panic(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	listener, err := tls.Listen("tcp", localTLS, config)
	if err != nil {
		if callback_error != nil {
			callback_error(err)
		}
		panic(err)
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

		go proxyConnection2TLS(conn, remoteTlsServerIp, remoteTlsPort, callback_error)
	}
}
