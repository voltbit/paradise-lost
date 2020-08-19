package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type MemdbAPIServer struct {
	addr               string
	peristenceFileName string
	logFileName        string
	persistenceFile    *os.File
	logFile            *os.File
	data               map[string]int64
}

func reloadData() error {
	// TODO: use mmap here for fun
	return nil
}

func NewMemdbAPIServer(url string, persistenceFileName string, logFileName string) (*MemdbAPIServer, error) {
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
	serverObj := &MemdbAPIServer{
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
	// swagger:route GET /
	// Main landing page
	fmt.Fprint(w, "Home page of the in-memory database")
	InfoLogger.Printf("[Request]Method: %s source: %s\n agent: %s", r.Method, r.RemoteAddr, r.UserAgent())
}

func (m *MemdbAPIServer) uploadText(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /api/v1/upload
	// Uploads text from the message body to the database
	// Responses:
	// 200: OK
	InfoLogger.Printf("[Request]Method: %s source: %s\n", r.Method, r.RemoteAddr)
	if r.Body == nil {
		DebugLogger.Println("no body received")
	}
	if r.Body != nil {
		scanner := bufio.NewScanner(r.Body)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			m.data[scanner.Text()]++
		}
		DebugLogger.Println("Data map after insertion", m.data)
	}
	w.Write([]byte("OK"))
}

func (m *MemdbAPIServer) getOccurence(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /api/v1/word
	// Get the number of occurences for a specified word
	// Produces:
	// - application/json
	// swagger:parameters word
	// min items: 1
	// in: query
	InfoLogger.Printf("[Request] Method: %s source: %s\n", r.Method, r.RemoteAddr)
	words := r.URL.Query()["word"]
	if len(words) < 1 {
		InfoLogger.Println("Empty query")
	}
	response := make(map[string]int64)
	for _, word := range words {
		response[word] = m.data[word]
	}
	DebugLogger.Println("Received query: ", words)
	jsonData, err := json.Marshal(response)
	if err != nil {
		ErrorLogger.Println("Could not load query data", err)
	}
	_, err = w.Write(jsonData)
	if err != nil {
		ErrorLogger.Println("Could not write data")
	}
}

func (m *MemdbAPIServer) getAllWords(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /api/v1/allwords
	// Dumps the whole database
	// Produces:
	// - application/json
	var res []byte
	var err error
	InfoLogger.Printf("[Request] Method: %s source: %s\n", r.Method, r.RemoteAddr)
	prettyFlag := r.URL.Query()["pretty"]
	if len(prettyFlag) > 0 {
		res, err = json.MarshalIndent(m.data, "", "    ")
	} else {
		res, err = json.Marshal(m.data)
	}
	if err != nil {
		// TODO handle this error
		ErrorLogger.Println("Could not get data", err)
	}
	w.Write(res)
}

func (m *MemdbAPIServer) Start() (err error) {
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
