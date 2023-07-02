package server

import (
	"os"

	"golang.org/x/exp/slog"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func SetupLogger() *slog.Logger {
	var logger *slog.Logger
	logLvl := os.Getenv("ENV")

	switch logLvl {
	case envDev:
		{
			logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		}
	}

	return logger
}
