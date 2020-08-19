package logger

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
)

func CreateLoggers(logFilePath string) error {
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file ", err)
		return err
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	DebugLogger = log.New(os.Stdout, "[DEBUG]", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(multiWriter, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(multiWriter, "[WARN]", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(multiWriter, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}
