package generos

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"
)

// Service encapsulates usecase logic for generos.
type Service interface {
	GetGeneros(ctx context.Context) ([]Generos, error)
	GetGeneroPorID(ctx context.Context, idGenero int) (Generos, error)
}

// Generos represents the data about an generos.
type Generos struct {
	entity.Genero
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new generos service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list generos.
func (s service) GetGeneros(ctx context.Context) ([]Generos, error) {
	generos, err := s.repo.GetGeneros(ctx)
	if err != nil {
		return nil, err
	}
	result := []Generos{}
	for _, item := range generos {
		result = append(result, Generos{item})
	}
	return result, nil
}

func (s service) GetGeneroPorID(ctx context.Context, idGenero int) (Generos, error) {
	genero, err := s.repo.GetGeneroPorID(ctx, idGenero)
	if err != nil {
		return Generos{}, err
	}
	return Generos{genero}, nil
}
