package infrastructure

import (
	"errors"
	"sort"
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

func (r *MemoryRepository) List(sortBy string, asc bool) ([]*domain.Exoplanet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	exoplanets := make([]*domain.Exoplanet, 0, len(r.data))
	for _, exoplanet := range r.data {
		exoplanets = append(exoplanets, exoplanet)
	}
	switch sortBy {
	case "radius":
		sort.Slice(exoplanets, func(i, j int) bool {
			if asc {
				return exoplanets[i].Radius < exoplanets[j].Radius
			}
			return exoplanets[i].Radius > exoplanets[j].Radius
		})
	case "mass":
		sort.Slice(exoplanets, func(i, j int) bool {
			if exoplanets[i].Mass == nil {
				return false
			}
			if exoplanets[j].Mass == nil {
				return true
			}
			if asc {
				return *exoplanets[i].Mass < *exoplanets[j].Mass
			}
			return *exoplanets[i].Mass > *exoplanets[j].Mass
		})
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
