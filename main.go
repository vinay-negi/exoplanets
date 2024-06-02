// main.go
package main

import (
	"log"
	"net/http"

	"github.com/vinay-negi/exoplanets/handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/exoplanets", handler.AddExoplanetHandler).Methods("POST")
	router.HandleFunc("/exoplanets", handler.ListExoplanetsHandler).Methods("GET")
	router.HandleFunc("/exoplanets/{id}", handler.GetExoplanetHandler).Methods("GET")
	router.HandleFunc("/exoplanets/{id}", handler.UpdateExoplanetHandler).Methods("PUT")
	router.HandleFunc("/exoplanets/{id}", handler.DeleteExoplanetHandler).Methods("DELETE")
	router.HandleFunc("/exoplanets/{id}/fuel-estimation", handler.FuelEstimationHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
