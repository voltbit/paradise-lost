package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type DummyToolConfig struct {
	LogGenConfig *LogGenConfig `yaml:logGen`
}

type LogGenConfig struct {
	TotalLogEntries int    `yaml:totalLogEntries`
	LogSize         string `yaml:logEntries`
}

func NewDummyToolConfig(path string) (*DummyToolConfig, error) {
	newConfig := new(DummyToolConfig)
	if err := newConfig.loadConfig(path); err != nil {
		return nil, err
	}
	return newConfig, nil
}

func (g *DummyToolConfig) loadConfig(path string) error {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		// r.Log.Error(err, "Failed to read the file")
		return err
	}
	if err := yaml.Unmarshal(buff, g); err != nil {
		return err
	}
	// r.Log.Info(fmt.Sprintf("Loaded config: %+v", g))
	return nil
}
