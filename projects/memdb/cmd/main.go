package main

import (
	"io"
	"log"
	"os"
)

var (
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
	InfoLogger = log.New(io.MultiWriter(os.Stdout, logFile), "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(io.MultiWriter(os.Stdout, logFile), "[WARN]", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(io.MultiWriter(os.Stdout, logFile), "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
	server, _ = NewMemdbServer("localhost:9889", "data_file", "log")
}

func main() {
	InfoLogger.Println("Initializing memstore")
	server.Start()
}
