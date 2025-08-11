package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"sso/internal/lib/logger"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting application", slog.String("env", cfg.Env))

	application := app.New(log, cfg.GRPC.Port)

	go application.GRPCServ.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("signal", sign.String()))

	application.GRPCServ.Stop()

	log.Info("application stopped")
}
