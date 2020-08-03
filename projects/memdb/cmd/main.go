//go:generate swagger generate spec
package main

import (
	"io"
	"log"
	"os"
)

var (
	DebugLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	server        *MemdbServer
)

func init() {
	logFile, err := os.OpenFile("log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file ", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	DebugLogger = log.New(os.Stdout, "[DEBUG]", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(multiWriter, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(multiWriter, "[WARN]", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(multiWriter, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
	server, _ = NewMemdbServer("localhost:9889", "data_file", "log")
}

// docs reference: https://goswagger.io/use/spec/route.html

func main() {
	InfoLogger.Println("Initializing memstore")
	server.Start()
}
