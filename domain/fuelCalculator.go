package domain

import "errors"

type FuelCalculator struct{}

func (fc *FuelCalculator) CalculateFuel(exoplanet *Exoplanet, crewCapacity int) (float64, error) {
	if crewCapacity <= 0 {
		return 0, errors.New("crew capacity must be greater than zero")
	}

	var gravity float64
	switch exoplanet.Type {
	case GasGiant:
		gravity = 0.5 / (exoplanet.Radius * exoplanet.Radius)
	case Terrestrial:
		gravity = *exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
	default:
		return 0, errors.New("unknown exoplanet type")
	}

	fuel := float64(exoplanet.Distance) / (gravity * gravity) * float64(crewCapacity)
	return fuel, nil
}
