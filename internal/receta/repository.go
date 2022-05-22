package receta

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access recetas from the data source.
type Repository interface {
	// GetRecetaPorId returns the receta with the specified receta ID.
	GetRecetaPorId(ctx context.Context, idReceta int) (entity.Receta, error)
	GetRecetaPorConsulta(ctx context.Context, idConsulta int) ([]entity.Receta, error)
	// GetRecetas returns the list recetas.
	GetRecetas(ctx context.Context) ([]entity.Receta, error)
	CrearReceta(ctx context.Context, receta entity.Receta) (entity.Receta, error)
	ActualizarReceta(ctx context.Context, receta entity.Receta) (entity.Receta, error)
}

// repository persists recetas in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new receta repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list recetas from the database.
func (r repository) GetRecetas(ctx context.Context) ([]entity.Receta, error) {
	var recetas []entity.Receta

	err := r.db.With(ctx).
		Select().
		From().
		OrderBy("apellidos asc").
		All(&recetas)
	if err != nil {
		return recetas, err
	}
	return recetas, err
}

// Create saves a new Receta record in the database.
// It returns the ID of the newly inserted receta record.
func (r repository) CrearReceta(ctx context.Context, receta entity.Receta) (entity.Receta, error) {
	err := r.db.With(ctx).Model(&receta).Insert()
	if err != nil {
		return entity.Receta{}, err
	}
	return receta, nil
}

// Create saves a new Receta record in the database.
// It returns the ID of the newly inserted receta record.
func (r repository) ActualizarReceta(ctx context.Context, receta entity.Receta) (entity.Receta, error) {
	var err error
	if receta.IdReceta != 0 {
		err = r.db.With(ctx).Model(&receta).Update()
	} else {
		err = r.db.With(ctx).Model(&receta).Insert()
	}
	if err != nil {
		return entity.Receta{}, err
	}
	return receta, nil
}

// Get reads the receta with the specified ID from the database.
func (r repository) GetRecetaPorId(ctx context.Context, idReceta int) (entity.Receta, error) {
	var receta entity.Receta
	err := r.db.With(ctx).Select().Model(idReceta, &receta)
	return receta, err
}

func (r repository) GetRecetaPorConsulta(ctx context.Context, idConsulta int) ([]entity.Receta, error) {
	var recetas []entity.Receta
	err := r.db.With(ctx).Select().Where(dbx.HashExp{"id_consulta": idConsulta}).All(&recetas)
	if err != nil {
		return recetas, err
	}
	return recetas, err
}
