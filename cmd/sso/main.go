package main

import (
	"log/slog"
	"sso/internal/config"
	"sso/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting application", slog.String("env", cfg.Env))
}
