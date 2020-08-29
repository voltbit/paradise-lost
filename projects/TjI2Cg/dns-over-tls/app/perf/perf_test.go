package perf

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"
	"testing"
)

func testProxy(b *testing.B, port string, proto string) {
	var tests = []struct {
		name string
	}{
		{"www.google.com"},
		{"www.slack.com"},
		{"www.github.com"},
		{"www.golang.org"},
	}
	for _, test := range tests {
		cmd := exec.Command("dig", "+noall", "+answer", test.name, "@localhost", "-p", port)
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()
		buf := new(bytes.Buffer)
		buf.ReadFrom(stdout)
		fmt.Println(buf.String())
	}
}

func benchmarkProxyUDP(b *testing.B, n int, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		testProxy(b, "11222", "udp")
	}
	wg.Done()
}

func benchmarkProxyTCP(b *testing.B, n int, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		testProxy(b, "11222", "tcp")
	}
	wg.Done()
}

func BenchmarkProxy(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(2)
	go benchmarkProxyUDP(b, 1, &wg)
	go benchmarkProxyTCP(b, 1, &wg)
	wg.Wait()
}
