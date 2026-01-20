package main

import (
	healthHttp "fiber-file-converter-api/internal/adapters/inbound/http/handler/health"
	metricsHttp "fiber-file-converter-api/internal/adapters/inbound/http/handler/metrics"
	"fiber-file-converter-api/internal/adapters/inbound/http/middleware"
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

	// middlewares
	// request id middleware
	app.Use(middleware.RequestID())
	// request logger middleware
	app.Use(middleware.RequestLogger())

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
