//go:generate swagger generate spec
package main

import node "../pkg/node"

func start_cluster(nodeCount int, config map[string]string) {
	for i := 0; i < nodeCount; i++ {
		node.NewMemdbNode("NODE" + string(i))
	}
}

// docs reference: https://goswagger.io/use/spec/route.html
func main() {
	standardConfig := make(map[string]string)
	standardConfig["logfile"] = ""
	standardConfig["persistencefile"] = ""
	standardConfig["nodelistfile"] = ""

	start_cluster(5, standardConfig)
}
