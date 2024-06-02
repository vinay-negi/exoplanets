package infrastructure

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/vinay-negi/exoplanets/application"
	"github.com/vinay-negi/exoplanets/domain"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *application.ExoplanetService
}

func NewHandler(service *application.ExoplanetService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) AddExoplanet(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string               `json:"name"`
		Description string               `json:"description"`
		Distance    int                  `json:"distance"`
		Radius      float64              `json:"radius"`
		Mass        *float64             `json:"mass,omitempty"`
		Type        domain.ExoplanetType `json:"type"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exoplanet, err := h.service.AddExoplanet(req.Name, req.Description, req.Distance, req.Radius, req.Mass, req.Type)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)
}

func (h *Handler) ListExoplanets(w http.ResponseWriter, r *http.Request) {
	exoplanets, err := h.service.ListExoplanets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(exoplanets)
}

func (h *Handler) GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	exoplanet, err := h.service.GetExoplanetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(exoplanet)
}

func (h *Handler) UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req struct {
		Name        string               `json:"name"`
		Description string               `json:"description"`
		Distance    int                  `json:"distance"`
		Radius      float64              `json:"radius"`
		Mass        *float64             `json:"mass,omitempty"`
		Type        domain.ExoplanetType `json:"type"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exoplanet, err := h.service.UpdateExoplanet(id, req.Name, req.Description, req.Distance, req.Radius, req.Mass, req.Type)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(exoplanet)
}

func (h *Handler) DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteExoplanet(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) EstimateFuel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	crewCapacityStr := r.URL.Query().Get("crewCapacity")
	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil {
		http.Error(w, "invalid crew capacity", http.StatusBadRequest)
		return
	}

	fuel, err := h.service.EstimateFuel(id, crewCapacity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]float64{"fuel_estimation": fuel})
}
