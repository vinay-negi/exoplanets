package infrastructure

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(handler *Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/exoplanets", handler.AddExoplanet).Methods(http.MethodPost)
	r.HandleFunc("/exoplanets", handler.ListExoplanets).Methods(http.MethodGet)
	r.HandleFunc("/exoplanets/{id}", handler.GetExoplanetByID).Methods(http.MethodGet)
	r.HandleFunc("/exoplanets/{id}", handler.UpdateExoplanet).Methods(http.MethodPut)
	r.HandleFunc("/exoplanets/{id}", handler.DeleteExoplanet).Methods(http.MethodDelete)
	r.HandleFunc("/exoplanets/{id}/fuel", handler.EstimateFuel).Methods(http.MethodGet)

	return r
}
