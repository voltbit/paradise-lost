package proxy

import (
	"sync"

	"./dnsserver"
	"./tlsserver"

	"golang.org/x/net/dns/dnsmessage"
)

func NewDNSTLSProxy() {
	var wg sync.WaitGroup

	wg.Add(2)
	proxyChannel := make(chan chan *dnsmessage.Message)
	go dnsserver.NewDNSServer(proxyChannel, &wg)
	go tlsserver.NewTLSServer(proxyChannel, &wg)
	wg.Wait()
}
