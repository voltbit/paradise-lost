package node

import (
	"math/rand"
	"sync"
	"time"

	api "../api"
	cfg "../cfg"
	cluster "../cluster"
	log "../logger"
)

var wg sync.WaitGroup

func generateId(length int) string {
	if length < 1 {
		length = 10
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	id := make([]rune, length)
	var letter rune
	for i := 0; i < length; i++ {
		if rand.Int()%2 == 0 {
			letter = 'a'
		} else {
			letter = 'A'
		}
		id[i] = letter + rune(rand.Intn(26))
	}
	return string(id)
}

func createClusterServer(nodeConfig cfg.NodeConfig, clusterConfig cfg.NodeConfigs) {
	url := nodeConfig.Address + ":" + nodeConfig.ClusterPort
	clusterComms := cluster.NewMemdbClusterComms(nodeConfig.NodeId, url)
	clusterComms.StartServer()
	clusterComms.StartPinging(2*time.Second, getListOfNodeAddr(clusterConfig))
}

func createAPIServer(nodeConfig cfg.NodeConfig) {
	url := nodeConfig.Address + ":" + nodeConfig.ApiPort
	apiServer, err := api.NewMemdbAPIServer(url, nodeConfig.PersistenceFile, nodeConfig.LogFile)
	if err != nil {
		log.ErrorLogger.Println("Could not create API server", err)
	}
	apiServer.Start()
}

func getListOfNodeAddr(clusterConfig cfg.NodeConfigs) [][]string {
	var nodes [][]string
	for _, node := range clusterConfig.Nodes {
		nodes = append(nodes, []string{node.NodeId, node.Address + ":" + node.ClusterPort})
	}
	return nodes
}

func CreateNode(nodeConfig cfg.NodeConfig, clusterConfig cfg.NodeConfigs) {
	wg.Add(2)
	go createAPIServer(nodeConfig)
	go createClusterServer(nodeConfig, clusterConfig)
	wg.Wait()
}
