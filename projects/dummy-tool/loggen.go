package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/bytefmt"
	"go.uber.org/zap"
)

type LogGen struct {
	Log    *zap.Logger
	Config *LogGenConfig
}

func NewLogGen(config *LogGenConfig) *LogGen {
	// mainLogger, err := zap.NewDevelopment()
	mainLogger, err := zap.NewProduction()
	mainLogger = mainLogger.Named("sre-dummy")
	if err != nil {
		fmt.Printf("Failed to create new logger")
		os.Exit(1)
	}
	mainLogger.Info(fmt.Sprintf("Test logs from the dummy app"))
	return &LogGen{
		Log: mainLogger,
	}
}

func (g *LogGen) GenerateLogBySize(size string) error {
	logByteSize, err := bytefmt.ToBytes(size)
	if err != nil {
		return err
	}
	fmt.Printf("Generating logs of size: %d\n", logByteSize)
	return nil
}
