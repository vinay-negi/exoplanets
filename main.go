package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vinay-negi/exoplanets/application"
	"github.com/vinay-negi/exoplanets/infrastructure"
)

func NewRouter(handler *infrastructure.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/exoplanets", handler.AddExoplanet).Methods(http.MethodPost)
	r.HandleFunc("/exoplanets", handler.ListExoplanets).Methods(http.MethodGet)
	r.HandleFunc("/exoplanets/{id}", handler.GetExoplanetByID).Methods(http.MethodGet)
	r.HandleFunc("/exoplanets/{id}", handler.UpdateExoplanet).Methods(http.MethodPut)
	r.HandleFunc("/exoplanets/{id}", handler.DeleteExoplanet).Methods(http.MethodDelete)
	r.HandleFunc("/exoplanets/{id}/fuel", handler.EstimateFuel).Methods(http.MethodGet)

	return r
}

func StartServer() {
	repo := infrastructure.NewMemoryRepository()
	service := application.NewExoplanetService(repo)
	handler := infrastructure.NewHandler(service)
	router := NewRouter(handler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}

func main() {
	StartServer()
}
