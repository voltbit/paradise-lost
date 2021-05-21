package main

import (
	"fmt"
	"os"
)

func main() {
	lines := 100
	token_len := 7
	err := generate_file(lines, token_len)
	if err != nil {
		fmt.Printf("Failed to run file generator %v\n", err)
		os.Exit(1)
	}
	processor, err := NewTokenProcessor("tokenuser", "tokenpass", "tokendb")
	if err != nil {
		fmt.Printf("Failed to create token processor: %v\n", err)
		os.Exit(1)
	}
	if err := processor.start(); err != nil {
		fmt.Printf("Failed to start processor: %v\n", err)
		os.Exit(1)
	}
}
