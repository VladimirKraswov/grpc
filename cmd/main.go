package main

import (
	"grpc/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	application := app.New(8080)

	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<- stop
	
	application.GRPCServer.Stop()
}
