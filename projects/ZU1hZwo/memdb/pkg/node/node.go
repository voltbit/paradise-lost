//go:generate swagger generate spec
package main

import (
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

const ID_LENGTH = 8

var (
	DebugLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	node          *MemdbNode
)

type MemdbNode struct {
	id        string
	apiServer *MemdbAPIServer
	p2pServer *MemdbP2PServer
}

func init_logging() {
	logFile, err := os.OpenFile("log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file ", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	DebugLogger = log.New(os.Stdout, "[DEBUG]", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(multiWriter, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(multiWriter, "[WARN]", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(multiWriter, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
}

func init() {
	init_logging()
}

func generateId(length int) string {
	if length < 1 {
		length = 10
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	id := make([]rune, length)
	var letter rune
	for i := 0; i < length; i++ {
		if rand.Int()%2 == 0 {
			letter = 'a'
		} else {
			letter = 'A'
		}
		id[i] = letter + rune(rand.Intn(26))
	}
	return string(id)
}

func NewMemdbNode(id string, addr string, config map[string]string) (*MemdbNode, error) {
	// TODO: create a config object based on a file to load all the relevant details
	apiServer, err := NewMemdbAPIServer(addr, config["persistencefile"], config["logfile"])
	if err != nil {
		ErrorLogger.Println("Memdb node could not initialize API server on node.", err)
		return nil, err
	}
	p2pServer, err := NewMemdbP2PServer()
	if err != nil {
		ErrorLogger.Println("Memdb node could not initialize P2P server.", err)
		return nil, err
	}
	if id == "" {
		id = getRandId(ID_LENGTH)
	}
	return &MemdbNode{
		id:        id,
		apiServer: apiServer,
		p2pServer: p2pServer,
	}, nil
}

func p2p_server() {

}

func Start() {
	node.apiServer.Start()
	node.p2pServer.Start()
}
