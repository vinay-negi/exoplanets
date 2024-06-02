package application

import (
	"github.com/vinay-negi/exoplanets/domain"
)

type ExoplanetService struct {
	repo Repository
	fc   *domain.FuelCalculator
}

type Repository interface {
	Add(exoplanet *domain.Exoplanet) error
	List(so string, asc bool) ([]*domain.Exoplanet, error)
	GetByID(id string) (*domain.Exoplanet, error)
	Update(exoplanet *domain.Exoplanet) error
	Delete(id string) error
}

func NewExoplanetService(repo Repository) *ExoplanetService {
	return &ExoplanetService{repo: repo, fc: &domain.FuelCalculator{}}
}

func (s *ExoplanetService) AddExoplanet(name, description string, distance int, radius float64, mass *float64, planetType domain.ExoplanetType) (*domain.Exoplanet, error) {
	exoplanet, err := domain.NewExoplanet(name, description, distance, radius, mass, planetType)
	if err != nil {
		return nil, err
	}
	err = s.repo.Add(exoplanet)
	return exoplanet, err
}

func (s *ExoplanetService) ListExoplanets(so string, asc bool) ([]*domain.Exoplanet, error) {
	return s.repo.List(so, asc)
}

func (s *ExoplanetService) GetExoplanetByID(id string) (*domain.Exoplanet, error) {
	return s.repo.GetByID(id)
}

func (s *ExoplanetService) UpdateExoplanet(id, name, description string, distance int, radius float64, mass *float64, planetType domain.ExoplanetType) (*domain.Exoplanet, error) {
	exoplanet, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	exoplanet.Name = name
	exoplanet.Description = description
	exoplanet.Distance = distance
	exoplanet.Radius = radius
	exoplanet.Mass = mass
	exoplanet.Type = planetType
	err = exoplanet.Validate()
	if err != nil {
		return nil, err
	}
	err = s.repo.Update(exoplanet)
	return exoplanet, err
}

func (s *ExoplanetService) DeleteExoplanet(id string) error {
	return s.repo.Delete(id)
}

func (s *ExoplanetService) EstimateFuel(id string, crewCapacity int) (float64, error) {
	exoplanet, err := s.repo.GetByID(id)
	if err != nil {
		return 0, err
	}
	return s.fc.CalculateFuel(exoplanet, crewCapacity)
}
