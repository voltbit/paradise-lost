package dnsserver

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"

	"../tools"

	"golang.org/x/net/dns/dnsmessage"
)

func manageUDPDNSRequest(buf []byte, addr net.Addr, conn net.PacketConn, requestCh chan *dnsmessage.Message) {
	dnsMsg := dnsmessage.Message{}
	err := dnsMsg.Unpack(buf)
	tools.CheckError("Could not unpack message", err)
	requestCh <- &dnsMsg
	response := <-requestCh
	buf = make([]byte, 512)
	buf, err = response.Pack()
	tools.CheckError("Could not pack data", err)
	fmt.Println("[DNS UDP Server] received answer:")
	tools.ShowPackage(response)
	conn.WriteTo(buf, addr)
}

func manageTCPDNSRequest(conn net.Conn, requestCh chan *dnsmessage.Message) {
	var err error

	dnsMsg := dnsmessage.Message{}
	buf := make([]byte, 512)
	_, err = conn.Read(buf)
	tools.CheckError("Could not read TCP message", err)
	err = dnsMsg.Unpack(buf[2:])
	tools.CheckError("Could not unpack message", err)
	fmt.Println("[DNS TCP Server] sending request for", dnsMsg.Questions)
	requestCh <- &dnsMsg
	response := <-requestCh
	fmt.Println("Received response:")
	tools.ShowPackage(response)
	buf, err = response.Pack()
	tools.CheckError("Could not pack data", err)
	buf = append(make([]byte, 2), buf...)
	binary.BigEndian.PutUint16(buf[:2], uint16(len(buf))-2)
	_, err = conn.Write(buf)
	tools.CheckError("Could not write response", err)
}

func listenUDP(proxyChannel chan chan *dnsmessage.Message, wg *sync.WaitGroup) {
	var err error
	var conn net.PacketConn
	var addr net.Addr

	defer wg.Done()

	conn, err = net.ListenPacket("udp", ":53")
	tools.CheckError("Connection error", err)
	defer conn.Close()
	fmt.Println("Listening for DNS messages")
	for {
		buf := make([]byte, 512)
		_, addr, err = conn.ReadFrom(buf)
		tools.CheckError("Read error in UDP DNS server", err)
		requestCh := make(chan *dnsmessage.Message)
		proxyChannel <- requestCh
		go manageUDPDNSRequest(buf, addr, conn, requestCh)
	}
}

func listenTCP(proxyChannel chan chan *dnsmessage.Message, wg *sync.WaitGroup) {
	var err error
	var ln net.Listener
	var conn net.Conn

	defer wg.Done()

	ln, err = net.Listen("tcp", ":53")
	tools.CheckError("Could not listen on TCP 53 port", err)

	for {
		conn, err = ln.Accept()
		tools.CheckError("Could not accept TCP connection", err)
		requestCh := make(chan *dnsmessage.Message)
		proxyChannel <- requestCh
		go manageTCPDNSRequest(conn, requestCh)
	}
}

func NewDNSServer(proxyChannel chan chan *dnsmessage.Message, mainWg *sync.WaitGroup) {
	var wg sync.WaitGroup
	wg.Add(2)
	go listenUDP(proxyChannel, &wg)
	go listenTCP(proxyChannel, &wg)
	wg.Wait()
	mainWg.Done()
}
