package accesos

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
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
	var idRol int

	err := r.db.With(ctx).
		Select("id_rol").
		From("usuario_rol").
		Where(dbx.HashExp{"id_usuario": idUsuario}).
		Row(&idRol)
	if err != nil {
		return accesos, 0, err
	}

	err = r.db.With(ctx).
		Select().
		From("accesos as a").
		InnerJoin("rol_acceso as ra", dbx.NewExp("ra.id_acceso = a.id_acceso")).
		Where(dbx.HashExp{"ra.id_rol": idRol}).
		OrderBy("a.ruta asc").
		All(&accesos)
	if err != nil {
		return accesos, 0, err
	}
	return accesos, idRol, err
}
