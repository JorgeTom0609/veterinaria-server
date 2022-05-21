package rol

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access roles from the data source.
type Repository interface {
	// GetRolPorId returns the rol with the specified rol ID.
	GetRolPorId(ctx context.Context, idRol int) (entity.Rol, error)
	// GetRoles returns the list roles.
	GetRoles(ctx context.Context) ([]entity.Rol, error)
	CrearRol(ctx context.Context, rol entity.Rol) (entity.Rol, error)
	ActualizarRol(ctx context.Context, rol entity.Rol) (entity.Rol, error)
}

// repository persists roles in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new rol repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list roles from the database.
func (r repository) GetRoles(ctx context.Context) ([]entity.Rol, error) {
	var roles []entity.Rol

	err := r.db.With(ctx).
		Select().
		From().
		All(&roles)
	if err != nil {
		return roles, err
	}
	return roles, err
}

// Create saves a new Rol record in the database.
// It returns the ID of the newly inserted rol record.
func (r repository) CrearRol(ctx context.Context, rol entity.Rol) (entity.Rol, error) {
	err := r.db.With(ctx).Model(&rol).Insert()
	if err != nil {
		return entity.Rol{}, err
	}
	return rol, nil
}

// Create saves a new Rol record in the database.
// It returns the ID of the newly inserted rol record.
func (r repository) ActualizarRol(ctx context.Context, rol entity.Rol) (entity.Rol, error) {
	var err error
	if rol.IdRol != 0 {
		err = r.db.With(ctx).Model(&rol).Update()
	} else {
		err = r.db.With(ctx).Model(&rol).Insert()
	}
	if err != nil {
		return entity.Rol{}, err
	}
	return rol, nil
}

// Get reads the rol with the specified ID from the database.
func (r repository) GetRolPorId(ctx context.Context, idRol int) (entity.Rol, error) {
	var rol entity.Rol
	err := r.db.With(ctx).Select().Model(idRol, &rol)
	return rol, err
}
