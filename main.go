package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		logger.Fatal("failed to create zap logger")
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info("This is importer-source")
}
