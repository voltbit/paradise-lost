package listener

import (
	"fmt"
	"net"
)

func checkError(s string, e error) {
	if e != nil {
		fmt.Println(s, e)
	}
}

func listenUDP() {
	var err error
	var conn net.PacketConn
	// conn, err = net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4("10.0.0.1"), Port: 53})
	conn, err = net.ListenPacket("udp", ":53")
	checkError("Connection error", err)
	defer conn.Close()
	fmt.Println("Listening for DNS messages")
	for {
		checkError("Failed to accept connection", err)
		buf := make([]byte, 512)
		n, addr, _ := conn.ReadFrom(buf)
		fmt.Println(n, addr)
		fmt.Println(buf)
	}
}

// func switchNetworkNamespace() {
// 	childNS, err := unix.Open("/run/netns/blue", unix.O_RDONLY, 0)
// 	checkError("Could not open new namespace", err)
// 	err = unix.Setns(childNS, unix.CLONE_NEWNET)
// 	checkError("Could not set namespace", err)
// }

func NewDNSListener() {
	// switchNetworkNamespace()
	listenUDP()
}
