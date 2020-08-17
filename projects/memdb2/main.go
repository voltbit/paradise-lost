package main

import (
	"fmt"
	"sync"

	cfg "./cfg"
	log "./logger"
	node "./node"
)

var wg sync.WaitGroup

// Creating a full local cluster based on provided configuration
func createCluster(c cfg.NodeConfigs) {
	for _, conf := range c.Nodes {
		wg.Add(1)
		log.InfoLogger.Println("Creating node", conf.NodeId, conf.ApiPort)
		go node.CreateNode(conf, c)
	}
}

func loadLoggers() {
	log.CreateLoggers("/tmp/memdb/logging")
}

func loadConfig(configFilePath string) cfg.NodeConfigs {
	return cfg.LoadConfig(configFilePath)
}

func main() {
	loadLoggers()
	conf := loadConfig("/home/andrei/development/paradise-lost/projects/memdb2/cluster.json")
	// conf := loadConfig("/home/andrei/development/paradise-lost/projects/memdb2/singleNode.json")
	log.InfoLogger.Println("Cluster controller started")
	fmt.Println("Got config:", conf)
	fmt.Println(conf.Nodes[0].ClusterPort)
	createCluster(conf)
	wg.Wait()
}
