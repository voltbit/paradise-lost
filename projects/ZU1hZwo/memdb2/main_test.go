package main

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config := loadConfig("/home/andrei/development/paradise-lost/projects/memdb2/cluster.json")
	fmt.Println(config)
	if len(config.Nodes) < 1 {
		t.Error("No node inforamtion found")
	}
	for _, nodeConf := range config.Nodes {
		if nodeConf.NodeId == "" {
			t.Error("Node has no ID")
		}
	}
}
