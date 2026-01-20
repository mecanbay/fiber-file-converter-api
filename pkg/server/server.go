package server

import (
	"fiber-file-converter-api/pkg/config"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var (
	MsgServerStarting   = "server starting..."
	MsgServerStarted    = "server started on port %v"
	MsgGracefulShutdown = "server gracefully stopped!"
	ErrServerStart      = "failed to start server! Err : %v"
	ErrGracefulShutdown = "error during server shutdown! Err : %v"
)

func New(cfg *config.AppConfig) *fiber.App {
	config := fiber.Config{
		AppName:               cfg.AppName,
		IdleTimeout:           cfg.IdleTimeout * time.Second,
		ReadTimeout:           cfg.ReadTimeout * time.Second,
		WriteTimeout:          cfg.WriteTimeout * time.Second,
		DisableStartupMessage: false,
		ErrorHandler:          errorHandler,
	}

	app := fiber.New(config)
	return app
}

func Start(app *fiber.App, cfg *config.AppConfig) {
	zap.L().Info(MsgServerStarting)
	addr := fmt.Sprintf(":%v", cfg.Port)

	go func() {
		if err := app.Listen(addr); err != nil {
			zap.L().Fatal(fmt.Sprintf(ErrServerStart, err.Error()))
		}
	}()
	zap.L().Info(fmt.Sprintf(MsgServerStarted, cfg.Port))
	gracefulShutdown(app, cfg.GracefulShutdownTimeout)
}

func gracefulShutdown(app *fiber.App, timeout time.Duration) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	if err := app.ShutdownWithTimeout(timeout * time.Second); err != nil {
		zap.L().Error(fmt.Sprintf(ErrGracefulShutdown, err.Error()))
	}
	zap.L().Info(MsgGracefulShutdown)
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	path := c.Path()

	if strings.HasPrefix(path, "/.well-known/") || path == "/favicon.icon" {
		return c.SendStatus(code)
	}

	fields := []zap.Field{
		zap.Error(err),
		zap.Int("status", code),
		zap.String("path", path),
		zap.String("method", c.Method()),
	}

	switch {
	case code >= 500:
		zap.L().Error("unhandled error", fields...)
	case code >= 400:
		zap.L().Warn("http client error", fields...)
	default:
		zap.L().Debug("http error", fields...)
	}

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})

}
