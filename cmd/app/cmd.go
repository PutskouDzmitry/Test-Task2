package main

import (
	"Test-Task2/internal/app"
	"Test-Task2/internal/logger"
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var log *zap.Logger

func init() {
	log = logger.CreateLogger()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := app.NewApp(ctx, log)
	if err := application.Run(); err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("Started application")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Debug("shutting down server...")

	log.Info("Stopped application")
}
