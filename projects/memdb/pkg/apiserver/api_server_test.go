package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestUploadFile(t *testing.T) {
	testTable := []struct {
		name     string
		filename string
	}{
		{
			name:     "GenericText",
			filename: "testdata/lorem_ipsum.txt",
		},
	}

	for _, testData := range testTable {
		t.Run(testData.name, func(t *testing.T) {
			fh, err := os.OpenFile(testData.filename, os.O_RDONLY, 0666)
			if err != nil {
				fmt.Println(err)
			}
			http.Post("localhost:9889", "", fh)
		})
	}
}

func TestGetWord(t *testing.T) {
	// http://localhost:9889/api/v1/wordcount?word=cubilia&word=eget
	// testTable = []struct {
	// 	input    string
	// 	expected string
	// }{}
	// for _, testinput := range testTable {

	// }
}
