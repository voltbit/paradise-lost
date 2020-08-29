package dnsserver

import (
	"fmt"
	"net"
	"sync"

	"../tools"

	"golang.org/x/net/dns/dnsmessage"
)

func manageUDPDNSRequest(buf []byte, addr net.Addr, requestCh chan *dnsmessage.Message) {
	dnsMsg := dnsmessage.Message{}
	err := dnsMsg.Unpack(buf)
	tools.CheckError("Could not unpack message", err)
	fmt.Println(dnsMsg.GoString())
	requestCh <- &dnsMsg
	response := <-requestCh
	fmt.Println("[DNS UDP Server] received answer", response.Answers)
}

func manageTCPDNSRequest(conn net.Conn) {
	var n int
	var err error

	dnsMsg := dnsmessage.Message{}
	buf := make([]byte, 512)
	n, err = conn.Read(buf)
	tools.CheckError("Could not read TCP message", err)
	fmt.Println("Size received", n)
	err = dnsMsg.Unpack(buf[2:])
	tools.CheckError("Could not unpack message", err)
	fmt.Println(dnsMsg.GoString())
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
		go manageUDPDNSRequest(buf, addr, requestCh)
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
		go manageTCPDNSRequest(conn)
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
