package listener

import (
	"fmt"
	"net"

	"golang.org/x/net/dns/dnsmessage"
)

func checkError(s string, e error) {
	if e != nil {
		fmt.Println(s, e)
	}
}

func listenUDP() {
	var err error
	var m dnsmessage.Message
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 53})
	checkError("Could not connect to port 53", err)
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		_, addr, _ := conn.ReadFromUDP(buf)
		err = m.Unpack(buf)
		checkError("Could not unpack message", err)
		fmt.Println("Received package from", addr)
		fmt.Println(buf)
	}
}

func NewDNSListener(p string) {
	listenUDP()
}
