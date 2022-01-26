package especies

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access especies from the data source.
type Repository interface {
	// Get returns the list especies.
	GetEspecies(ctx context.Context) ([]entity.Especie, error)
}

// repository persists especies in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new especie repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list especies from the database.
func (r repository) GetEspecies(ctx context.Context) ([]entity.Especie, error) {
	var especies []entity.Especie

	err := r.db.With(ctx).
		Select().
		From("especies").
		All(&especies)
	if err != nil {
		return especies, err
	}
	return especies, err
}
