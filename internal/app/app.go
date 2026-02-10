package app

import (
	"context"
	"log/slog"
	"scraper/internal/logger"
)

func Run(ctx context.Context) error {
	
	opts := logger.Options{
		Level: slog.LevelDebug,
		AddSource: false,
	}
	
	log := logger.New(opts)
	log.Info("app works!")


	return nil
}