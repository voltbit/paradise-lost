package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type MemdbServer struct {
	addr               string
	peristenceFileName string
	logFileName        string
	persistenceFile    *os.File
	logFile            *os.File
	data               map[string]int64
}

func reloadData() error {
	return nil
}

func NewMemdbServer(url string, persistenceFileName string, logFileName string) (*MemdbServer, error) {
	var err error
	if url == "" {
		url = "localhost:9889"
	}
	if logFileName == "" {
		logFileName = "log"
	}
	if persistenceFileName == "" {
		persistenceFileName = "data_file"
	}
	serverObj := &MemdbServer{
		addr:               url,
		peristenceFileName: persistenceFileName,
		logFileName:        logFileName,
		persistenceFile:    nil,
		logFile:            nil,
		data:               make(map[string]int64),
	}
	// TODO check to see what are the best permisions for these files
	serverObj.logFile, err = os.OpenFile(logFileName, os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		ErrorLogger.Println("Could not open log file", err)
		return nil, err
	}
	serverObj.persistenceFile, err = os.OpenFile(persistenceFileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		ErrorLogger.Println("Could not open log file", err)
		return nil, err
	}

	return serverObj, nil
}

func landPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page of the in-memory database")
	InfoLogger.Printf("[Request]Method: %s source: %s\n agent: %s", r.Method, r.RemoteAddr, r.UserAgent())
}

func (m *MemdbServer) uploadText(w http.ResponseWriter, r *http.Request) {
	InfoLogger.Printf("[Request]Method: %s source: %s\n", r.Method, r.RemoteAddr)
	InfoLogger.Println("Received text:\n", r.Body)
	if r.Body != nil {
		scanner := bufio.NewScanner(r.Body)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			m.data[scanner.Text()]++
		}
	}
}

func (m *MemdbServer) getOccurence(w http.ResponseWriter, r *http.Request) {
	InfoLogger.Printf("[Request]Method: %s source: %s\n", r.Method, r.RemoteAddr)
	jsonData, err := json.Marshal(m.data["test"])
	if err != nil {
		ErrorLogger.Println("Could not load query data", err)
	}
	_, err = w.Write(jsonData)
	if err != nil {
		ErrorLogger.Println("Could not write data")
	}
}

func (m *MemdbServer) getAllWords(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(m.data)
	if err != nil {
		// TODO handle this error
		ErrorLogger.Println("Could not get data", err)
	}
	w.Write(res)
}

// what happens if I call the method multiple times when receiving a value as struct target?
// if a new process is created then what happens with the old one
func (m *MemdbServer) Start() (err error) {
	InfoLogger.Println("Starting server")

	defer func() {
		if cleanupErr := m.persistenceFile.Sync(); cleanupErr != nil {
			ErrorLogger.Println(" !!! Persistent file data sync error", cleanupErr)
			err = cleanupErr
		}
		if cleanupErr := m.persistenceFile.Close(); cleanupErr != nil {
			ErrorLogger.Println("Persistent file data close error", cleanupErr)
			err = cleanupErr
		}
		if cleanupErr := m.logFile.Close(); cleanupErr != nil {
			ErrorLogger.Println("Log file close error", cleanupErr)
			err = cleanupErr
		}
	}()

	http.HandleFunc("/", landPage)
	http.HandleFunc("/api/v1/upload", m.uploadText)
	http.HandleFunc("/api/v1/wordcount", m.getOccurence)
	http.HandleFunc("/api/v1/allwords", m.getAllWords)

	err = http.ListenAndServe(m.addr, nil)
	if err != nil {
		ErrorLogger.Println("Failed to start the server", err)
		return
	}
	return
}
