package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HarshitTomar143/students-api/internal/config"
)

func main() {
	// Loading the config
	cfg := config.MustLoad()

	// Setting the router
	router := http.NewServeMux()

	router.HandleFunc("/Hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the student's api!")) // making a splice of byte data.
	})

	//setting the server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Printf("Server is running on %s", cfg.HTTPServer.Address)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start the server")
	}

}
