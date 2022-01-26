package generos

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access generos from the data source.
type Repository interface {
	// Get returns the list generos.
	GetGeneros(ctx context.Context) ([]entity.Genero, error)
}

// repository persists generos in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new genero repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list generos from the database.
func (r repository) GetGeneros(ctx context.Context) ([]entity.Genero, error) {
	var generos []entity.Genero

	err := r.db.With(ctx).
		Select().
		From("generos").
		All(&generos)
	if err != nil {
		return generos, err
	}
	return generos, err
}
