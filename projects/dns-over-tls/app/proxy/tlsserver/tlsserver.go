package tlsserver

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"os"
	"sync"

	"../tools"

	"golang.org/x/net/dns/dnsmessage"
)

func handleRequest(requestCh chan *dnsmessage.Message) {
	var logfd *os.File
	msg := <-requestCh
	var reply dnsmessage.Message
	fmt.Println("[TLS Server] received request for", msg.Questions)

	roots, err := x509.SystemCertPool()
	tools.CheckError("Could not make the certificates pool", err)
	logfd, err = os.OpenFile("/app/ssl-key.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	tools.CheckError("Could not open pre-master secret log file", err)

	defer logfd.Close()
	conn, err := tls.Dial("tcp", "1.1.1.1:853", &tls.Config{
		RootCAs:      roots,
		ServerName:   "cloudflare-dns.com",
		KeyLogWriter: logfd,
	})
	tools.CheckError("Could not connect to TLS server", err)
	defer conn.Close()

	state := conn.ConnectionState()
	fmt.Println("Remote DNS server handshake complete:", state.HandshakeComplete)

	buf := make([]byte, 512)
	buf, err = msg.Pack()
	tools.CheckError("Could not pack request", err)
	buf = append(make([]byte, 2), buf...)
	binary.BigEndian.PutUint16(buf[:2], uint16(len(buf))-2)

	fmt.Println("========================================")
	_, err = conn.Write(buf)
	tools.CheckError("Could not write buffer", err)

	rbuf := make([]byte, 512)
	_, err = conn.Read(rbuf)
	tools.CheckError("Could not read data", err)
	err = reply.Unpack(rbuf[2:])
	tools.CheckError("Could not unpack data", err)
	requestCh <- &reply
}

func listenRequests(proxyChannel chan chan *dnsmessage.Message, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		req := <-proxyChannel
		go handleRequest(req)
	}
}

func NewTLSServer(proxyChannel chan chan *dnsmessage.Message, mainWg *sync.WaitGroup) {
	var wg sync.WaitGroup
	wg.Add(1)
	go listenRequests(proxyChannel, &wg)
	wg.Wait()
	mainWg.Done()
}
