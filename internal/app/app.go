package app

import (
	grpcapp "grpc/internal/app/grpc"
	"grpc/internal/services/auth"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(grpcPort int) *App{
	authService := auth.New()
	grpcApp := grpcapp.New(authService, grpcPort)
	
	return &App{
		GRPCServer: grpcApp,
	}
}