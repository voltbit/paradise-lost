package tools

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/dns/dnsmessage"
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

func ShowPackage(msg *dnsmessage.Message) {
	fmt.Println("ID:", msg.Header.ID)
	fmt.Println("Response:", msg.Header.Response)
	fmt.Println("Questions:", msg.Questions)
	fmt.Println("Answers:", msg.Answers)
}
