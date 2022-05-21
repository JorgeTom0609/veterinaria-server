package usuario_rol

import (
	"context"
	"database/sql"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access usuarioRoles from the data source.
type Repository interface {
	// GetUsuarioRolPorId returns the usuarioRol with the specified usuarioRol ID.
	GetUsuarioRolPorId(ctx context.Context, idUsuarioRol int) (entity.UsuarioRol, error)
	GetUsuarioRolPorCedula(ctx context.Context, cedula string) (entity.UsuarioRol, error)
	// GetUsuarioRoles returns the list usuarioRoles.
	GetUsuarioRoles(ctx context.Context) ([]entity.UsuarioRol, error)
	CrearUsuarioRol(ctx context.Context, usuarioRol entity.UsuarioRol) (entity.UsuarioRol, error)
	ActualizarUsuarioRol(ctx context.Context, usuarioRol entity.UsuarioRol) (entity.UsuarioRol, error)
}

// repository persists usuarioRoles in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new usuarioRol repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list usuarioRoles from the database.
func (r repository) GetUsuarioRoles(ctx context.Context) ([]entity.UsuarioRol, error) {
	var usuarioRoles []entity.UsuarioRol

	err := r.db.With(ctx).
		Select().
		From().
		All(&usuarioRoles)
	if err != nil {
		return usuarioRoles, err
	}
	return usuarioRoles, err
}

// Create saves a new UsuarioRol record in the database.
// It returns the ID of the newly inserted usuarioRol record.
func (r repository) CrearUsuarioRol(ctx context.Context, usuarioRol entity.UsuarioRol) (entity.UsuarioRol, error) {
	err := r.db.With(ctx).Model(&usuarioRol).Insert()
	if err != nil {
		return entity.UsuarioRol{}, err
	}
	return usuarioRol, nil
}

// Create saves a new UsuarioRol record in the database.
// It returns the ID of the newly inserted usuarioRol record.
func (r repository) ActualizarUsuarioRol(ctx context.Context, usuarioRol entity.UsuarioRol) (entity.UsuarioRol, error) {
	var err error
	if usuarioRol.IdUsuarioRol != 0 {
		err = r.db.With(ctx).Model(&usuarioRol).Update()
	} else {
		err = r.db.With(ctx).Model(&usuarioRol).Insert()
	}
	if err != nil {
		return entity.UsuarioRol{}, err
	}
	return usuarioRol, nil
}

// Get reads the usuarioRol with the specified ID from the database.
func (r repository) GetUsuarioRolPorId(ctx context.Context, idUsuarioRol int) (entity.UsuarioRol, error) {
	var usuarioRol entity.UsuarioRol
	err := r.db.With(ctx).Select().Model(idUsuarioRol, &usuarioRol)
	return usuarioRol, err
}

func (r repository) GetUsuarioRolPorCedula(ctx context.Context, cedula string) (entity.UsuarioRol, error) {
	var usuarioRol entity.UsuarioRol
	err := r.db.With(ctx).Select().Where(dbx.HashExp{"cedula": cedula}).One(&usuarioRol)
	if err == sql.ErrNoRows {
		err = nil
	}
	return usuarioRol, err
}
