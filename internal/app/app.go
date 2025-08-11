package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
)

type App struct {
	GRPCServ *grpcapp.App
}

func New(log *slog.Logger, gRPCPort int) *App {
	grpcApp := grpcapp.New(log, gRPCPort)

	return &App{
		GRPCServ: grpcApp,
	}
}
