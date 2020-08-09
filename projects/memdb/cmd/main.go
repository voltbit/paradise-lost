//go:generate swagger generate spec
package main

// docs reference: https://goswagger.io/use/spec/route.html
func main() {
	node, _ = NewMemdbNode("")
	node.Start()
}
