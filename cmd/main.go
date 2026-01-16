package main

import (
	"fiber-file-converter-api/pkg/config"
	"fiber-file-converter-api/pkg/log"
)

func main() {
	// load config file
	cfg := config.Load()

	// Initalize zap logger
	logger := log.Load(cfg.App.Env, cfg.App.LogPath, cfg.App.LogFile)
	defer logger.Sync()
}
