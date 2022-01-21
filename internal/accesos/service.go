package accesos

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"
)

// Service encapsulates usecase logic for accesos.
type Service interface {
	GetAccesosPorIdUsuario(ctx context.Context, idUsuario int) ([]Accesos, int, error)
}

// Accesos represents the data about an accesos.
type Accesos struct {
	entity.Acceso
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new accesos service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list accesos with the specified the IdUsuario.
func (s service) GetAccesosPorIdUsuario(ctx context.Context, idUsuario int) ([]Accesos, int, error) {
	accesos, idRol, err := s.repo.GetAccesosPorIdUsuario(ctx, idUsuario)
	if err != nil {
		return nil, 0, err
	}
	result := []Accesos{}
	for _, item := range accesos {
		result = append(result, Accesos{item})
	}
	return result, idRol, nil
}
