// models.go
package models

import (
	"errors"
	"math"
	"sync"
)

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    int           `json:"distance"`       // in light years
	Radius      float64       `json:"radius"`         // in Earth-radius units
	Mass        float64       `json:"mass,omitempty"` // in Earth-mass units, only for Terrestrial
	Type        ExoplanetType `json:"type"`
}

type exoplanetStore struct {
	sync.Mutex
	Exoplanets map[string]Exoplanet
}

var Store = exoplanetStore{
	Exoplanets: make(map[string]Exoplanet),
}

func (e *Exoplanet) CalculateGravity() (float64, error) {
	switch e.Type {
	case GasGiant:
		return 0.5 / math.Pow(e.Radius, 2), nil
	case Terrestrial:
		if e.Mass == 0 {
			return 0, errors.New("mass is required for Terrestrial planets")
		}
		return e.Mass / math.Pow(e.Radius, 2), nil
	default:
		return 0, errors.New("unknown exoplanet type")
	}
}

func (e *Exoplanet) CalculateFuelEstimation(crewCapacity int) (float64, error) {
	gravity, err := e.CalculateGravity()
	if err != nil {
		return 0, err
	}
	return float64(e.Distance) / math.Pow(gravity, 2) * float64(crewCapacity), nil
}
