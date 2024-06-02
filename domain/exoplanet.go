package domain

import (
	"errors"

	"github.com/google/uuid"
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
	Mass        *float64      `json:"mass,omitempty"` // in Earth-mass units, only for Terrestrial
	Type        ExoplanetType `json:"type"`
}

func (e *Exoplanet) Validate() error {
	if e.Distance < 10 || e.Distance > 1000 ||
		e.Radius < 0.1 || e.Radius > 10 {
		return errors.New("invalid exoplanet data")
	}
	if e.Type == Terrestrial && (e.Mass == nil ||
		*e.Mass < 0.1 || *e.Mass > 10) {
		return errors.New("invalid mass for terrestrial exoplanet")
	}
	if e.Type != Terrestrial && e.Type != GasGiant {
		return errors.New("type must be either GasGiant or Terrestrial")
	}
	return nil
}

func NewExoplanet(name, description string, distance int, radius float64, mass *float64, planetType ExoplanetType) (*Exoplanet, error) {
	id := uuid.New().String()

	exoplanet := &Exoplanet{
		ID:          id,
		Name:        name,
		Description: description,
		Distance:    distance,
		Radius:      radius,
		Mass:        mass,
		Type:        planetType,
	}

	if err := exoplanet.Validate(); err != nil {
		return nil, err
	}

	return exoplanet, nil
}
