// validation.go
package applicaion

import (
	"errors"

	"github.com/vinay-negi/exoplanets/models"
)

func ValidateExoplanetInput(e *models.Exoplanet) error {
	if e.Distance <= 10 || e.Distance >= 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if e.Radius <= 0.1 || e.Radius >= 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if e.Type == "Terrestrial" && (e.Mass <= 0.1 || e.Mass >= 10) {
		return errors.New("mass must be between 0.1 and 10 Earth-mass units for Terrestrial planets")
	}
	if e.Type != "GasGiant" && e.Type != "Terrestrial" {
		return errors.New("type must be either GasGiant or Terrestrial")
	}
	return nil
}
