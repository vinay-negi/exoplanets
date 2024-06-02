package main

import (
	"log"
	"net/http"

	"github.com/vinay-negi/exoplanets/application"
	"github.com/vinay-negi/exoplanets/infrastructure"
)

func StartServer() {
	repo := infrastructure.NewMemoryRepository()
	service := application.NewExoplanetService(repo)
	handler := infrastructure.NewHandler(service)
	router := infrastructure.NewRouter(handler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}

func main() {
	StartServer()
}
