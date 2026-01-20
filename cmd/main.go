package main

import (
	healthHttp "fiber-file-converter-api/internal/adapters/inbound/http/handler/health"
	metricsHttp "fiber-file-converter-api/internal/adapters/inbound/http/handler/metrics"
	healthApp "fiber-file-converter-api/internal/application/health"
	"fiber-file-converter-api/internal/observability/metrics"
	"fiber-file-converter-api/pkg/config"
	"fiber-file-converter-api/pkg/log"
	"fiber-file-converter-api/pkg/server"
)

func main() {
	// load config file
	cfg := config.Load()

	// Initalize zap logger
	logger := log.Load(cfg.App.Env, cfg.App.LogPath, cfg.App.LogFile)
	defer logger.Sync()

	// Initialize metrics
	metrics.Init()

	// Initalize Fiber Server
	app := server.New(&cfg.App)

	// dependency injection

	healthService := healthApp.NewService()

	healthHandler := healthHttp.NewHandler(healthService)
	metricsHandler := metricsHttp.NewHandler()

	// setup routes
	healthHandler.SetupRoutes(app)
	metricsHandler.SetupRoutes(app)

	// start fiber http server
	server.Start(app, &cfg.App)

}
