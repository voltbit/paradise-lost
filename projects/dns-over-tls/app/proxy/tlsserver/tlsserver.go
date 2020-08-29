package tlsserver

import (
	"fmt"
	"sync"

	"golang.org/x/net/dns/dnsmessage"
)

func handleRequest(requestCh chan *dnsmessage.Message) {
	msg := <-requestCh
	fmt.Println("[TLS Server] received request for", msg.Questions)
	requestCh <- msg
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
