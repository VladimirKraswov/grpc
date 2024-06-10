package grpcapp

import (
	"fmt"
	authgrpc "grpc/internal/grpc/auth"
	"log"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port 			 int
}

func New(authService authgrpc.Auth, port int) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer, authService)

	return &App{
		gRPCServer: gRPCServer,
		port: port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}

}

func (a *App) Run() error {
	op := "grpcapp.Run"
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		fmt.Errorf(op, err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf(op, err)
	}

	log.Println("server running", err)

	return nil
}

func (a *App) Stop() {

}