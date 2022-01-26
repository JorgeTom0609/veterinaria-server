package especies

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"
)

// Service encapsulates usecase logic for especies.
type Service interface {
	GetEspecies(ctx context.Context) ([]Especies, error)
}

// Especies represents the data about an especies.
type Especies struct {
	entity.Especie
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new especies service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list especies.
func (s service) GetEspecies(ctx context.Context) ([]Especies, error) {
	especies, err := s.repo.GetEspecies(ctx)
	if err != nil {
		return nil, err
	}
	result := []Especies{}
	for _, item := range especies {
		result = append(result, Especies{item})
	}
	return result, nil
}
