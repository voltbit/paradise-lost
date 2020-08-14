package node

import (
	"math/rand"
	"sync"
	"time"

	api "../api"
	cfg "../cfg"
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

func createClusterServer() {

}

func createAPIServer(nodeConfig cfg.NodeConfig) {
	url := nodeConfig.Address + ":" + nodeConfig.ApiPort
	apiServer, err := api.NewMemdbAPIServer(url, nodeConfig.PersistenceFile, nodeConfig.LogFile)
	if err != nil {
		log.ErrorLogger.Println("Could not create API server", err)
	}
	apiServer.Start()
}

func CreateNode(nodeConfig cfg.NodeConfig) {
	wg.Add(2)
	go createAPIServer(nodeConfig)
	wg.Wait()
}
