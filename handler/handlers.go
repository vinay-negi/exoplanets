package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	applicaion "github.com/vinay-negi/exoplanets/application"
	"github.com/vinay-negi/exoplanets/models"
)

func AddExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	var e models.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := applicaion.ValidateExoplanetInput(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	e.ID = uuid.New().String()
	models.Store.Lock()
	models.Store.Exoplanets[e.ID] = e
	models.Store.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func ListExoplanetsHandler(w http.ResponseWriter, r *http.Request) {
	models.Store.Lock()
	var exoplanets []models.Exoplanet
	for _, exoplanet := range models.Store.Exoplanets {
		exoplanets = append(exoplanets, exoplanet)
	}
	models.Store.Unlock()
	json.NewEncoder(w).Encode(exoplanets)
}

func GetExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	models.Store.Lock()
	exoplanet, exists := models.Store.Exoplanets[id]
	models.Store.Unlock()
	if !exists {
		http.Error(w, "exoplanet not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(exoplanet)
}

func UpdateExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updatedExoplanet models.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&updatedExoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := applicaion.ValidateExoplanetInput(&updatedExoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.Store.Lock()
	_, exists := models.Store.Exoplanets[id]
	if !exists {
		models.Store.Unlock()
		http.Error(w, "exoplanet not found", http.StatusNotFound)
		return
	}
	updatedExoplanet.ID = id
	models.Store.Exoplanets[id] = updatedExoplanet
	models.Store.Unlock()
	json.NewEncoder(w).Encode(updatedExoplanet)
}

func DeleteExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	models.Store.Lock()
	_, exists := models.Store.Exoplanets[id]
	if exists {
		delete(models.Store.Exoplanets, id)
	}
	models.Store.Unlock()
	if !exists {
		http.Error(w, "exoplanet not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func FuelEstimationHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var request struct {
		CrewCapacity int `json:"crew_capacity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exoplanet, exists := models.Store.Exoplanets[id]
	if !exists {
		http.Error(w, "exoplanet not found", http.StatusNotFound)
		return
	}
	if request.CrewCapacity <= 0 {
		http.Error(w, "Invalid crew capacity", http.StatusBadRequest)
		return
	}
	fuel, err := exoplanet.CalculateFuelEstimation(request.CrewCapacity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]float64{"fuel_estimation": fuel})
}
