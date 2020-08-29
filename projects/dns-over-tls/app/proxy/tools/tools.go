package tools

import (
	"encoding/json"
	"fmt"
)

func CheckError(s string, e error) {
	if e != nil {
		fmt.Println("\033[31m"+s, e, "\033[0m")
	}
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
