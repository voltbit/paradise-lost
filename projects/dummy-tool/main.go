package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := NewDummyToolConfig("/usr/local/etc/dummy.conf")
	if err != nil {
		fmt.Printf("Failed to load config")
		os.Exit(1)
	}
	logGen := NewLogGen(config.LogGenConfig)
	logGen.GenerateLogBySize("33K")
}
