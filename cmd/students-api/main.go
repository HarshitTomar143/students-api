package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HarshitTomar143/students-api/internal/config"
	"github.com/HarshitTomar143/students-api/internal/http/handlers/student"
)

func main() {
	// Loading the config
	cfg := config.MustLoad()

	// Setting the router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	//setting the server
	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}
	slog.Info("server started",slog.String("address",cfg.HTTPServer.Address))
	fmt.Printf("Server is running on %s", cfg.HTTPServer.Address)

	done:= make(chan os.Signal, 1) // we are making a buffer channel

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// if the server is stopped then the data will be passed to the done channel and the blocking will be cleared
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start the server")
	}
	}()

	<-done // until we get the value in the channel this will be the blocking step

	slog.Info("shutting down the server") // logging the shutdown message

	// we are shutting down gracefully so that if the server is managing any request it can complete it before shutting down

	ctx,cancel:= context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err:= server.Shutdown(ctx)

	if err!= nil {
		slog.Error("failed to shutdown the server", slog.String("error",err.Error()))
	}

	slog.Info("server stopped successfully")
	

}
