package main

type MemdbClient struct {
	addr    string
	logFile string
}

func NewClient(addr string, logFile string) *MemdbClient {
	return &MemdbClient{
		addr:    addr,
		logFile: logFile,
	}
}
