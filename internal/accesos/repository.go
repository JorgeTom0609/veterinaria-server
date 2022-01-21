package accesos

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access accesos from the data source.
type Repository interface {
	// Get returns the list accesos with the specified IdUsuario.
	GetAccesosPorIdUsuario(ctx context.Context, idUsuario int) ([]entity.Acceso, int, error)
}

// repository persists accesos in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new acceso repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list accesos with the specified IdUsuario from the database.
func (r repository) GetAccesosPorIdUsuario(ctx context.Context, idUsuario int) ([]entity.Acceso, int, error) {
	var accesos []entity.Acceso
	//var rol entity.Rol
	err := r.db.With(ctx).
		Select().
		From("accesos").
		All(&accesos)
	return accesos, 0, err
}
