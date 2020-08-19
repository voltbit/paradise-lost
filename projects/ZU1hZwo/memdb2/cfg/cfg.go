package cfg

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "../logger"
)

type NodeConfigs struct {
	Nodes []NodeConfig `json:"nodes"`
}

type NodeConfig struct {
	NodeId          string `json:"nodeId"`
	Address         string `json:"address"`
	ApiPort         string `json:"apiPort"`
	ClusterPort     string `json:"clusterPort"`
	PersistenceFile string `json:"persistenceFile"`
	LogFile         string `json:"logFile"`
	Leader          bool   `json:"leader"`
}

func LoadConfig(path string) NodeConfigs {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.ErrorLogger.Println("Could not load")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var jsonConfig NodeConfigs
	json.Unmarshal(byteValue, &jsonConfig)

	return jsonConfig
}
