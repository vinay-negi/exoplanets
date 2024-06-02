package infrastructure

import (
	"errors"
	"sync"

	"github.com/vinay-negi/exoplanets/domain"
)

type MemoryRepository struct {
	data map[string]*domain.Exoplanet
	mu   sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{data: make(map[string]*domain.Exoplanet)}
}

func (r *MemoryRepository) Add(exoplanet *domain.Exoplanet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[exoplanet.ID] = exoplanet
	return nil
}

func (r *MemoryRepository) List() ([]*domain.Exoplanet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	exoplanets := make([]*domain.Exoplanet, 0, len(r.data))
	for _, exoplanet := range r.data {
		exoplanets = append(exoplanets, exoplanet)
	}
	return exoplanets, nil
}

func (r *MemoryRepository) GetByID(id string) (*domain.Exoplanet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	exoplanet, exists := r.data[id]
	if !exists {
		return nil, errors.New("exoplanet not found")
	}
	return exoplanet, nil
}

func (r *MemoryRepository) Update(exoplanet *domain.Exoplanet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[exoplanet.ID] = exoplanet
	return nil
}

func (r *MemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, id)
	return nil
}
