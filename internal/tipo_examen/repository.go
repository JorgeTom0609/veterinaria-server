package tipo_examen

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access tipoExamen from the data source.
type Repository interface {
	// GetTipoExamenPorId returns the tipoExamen with the specified tipoExamen ID.
	GetTipoExamenPorId(ctx context.Context, idTipoExamen int) (entity.TipoExamen, error)
	// GetTipoExamenes returns the list tipoExamenes.
	GetTipoExamenes(ctx context.Context) ([]entity.TipoExamen, error)
	GetTipoExamenPorEspecie(ctx context.Context, idEspecie int) ([]entity.TipoExamen, error)
	GetDetallesExamenPorTipoExamen(ctx context.Context, idTipoExamen int) (DetallesExamen, error)
	CrearTipoExamen(ctx context.Context, tipoExamen entity.TipoExamen) (entity.TipoExamen, error)
	ActualizarTipoExamen(ctx context.Context, tipoExamen entity.TipoExamen) (entity.TipoExamen, error)
}

// repository persists tipoExamenes in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new tipoExamen repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// GetTipoExamenes reads the list tipoExamenes from the database.
func (r repository) GetTipoExamenes(ctx context.Context) ([]entity.TipoExamen, error) {
	var tipoExamenes []entity.TipoExamen

	err := r.db.With(ctx).
		Select().
		From().
		All(&tipoExamenes)
	if err != nil {
		return tipoExamenes, err
	}
	return tipoExamenes, err
}

func (r repository) GetTipoExamenPorEspecie(ctx context.Context, idEspecie int) ([]entity.TipoExamen, error) {
	var tipoExamenes []entity.TipoExamen

	err := r.db.With(ctx).
		Select().
		From().
		Where(dbx.HashExp{"id_especie": idEspecie}).
		All(&tipoExamenes)
	if err != nil {
		return tipoExamenes, err
	}
	return tipoExamenes, err
}
func (r repository) GetDetallesExamenPorTipoExamen(ctx context.Context, idTipoExamen int) (DetallesExamen, error) {
	var detallesExamenCualitativo []entity.DetallesExamenCualitativo
	var detallesExamenCuantitativo []entity.DetallesExamenCuantitativo
	var detallesExamenInformativo []entity.DetallesExamenInformativo

	err := r.db.With(ctx).
		Select().
		From("detalles_examen_cualitativo").
		Where(dbx.HashExp{"id_tipo_examen": idTipoExamen}).
		All(&detallesExamenCualitativo)
	if err != nil {
		return DetallesExamen{}, err
	}
	err = r.db.With(ctx).
		Select().
		From("detalles_examen_cuantitativo").
		Where(dbx.HashExp{"id_tipo_examen": idTipoExamen}).
		All(&detallesExamenCuantitativo)
	if err != nil {
		return DetallesExamen{}, err
	}
	err = r.db.With(ctx).
		Select().
		From("detalles_examen_informativo").
		Where(dbx.HashExp{"id_tipo_examen": idTipoExamen}).
		All(&detallesExamenInformativo)
	if err != nil {
		return DetallesExamen{}, err
	}
	return DetallesExamen{DetallesExamenCualitativo: detallesExamenCualitativo, DetallesExamenCuantitativo: detallesExamenCuantitativo, DetallesExamenInformativo: detallesExamenInformativo}, err
}

// Create saves a new TipoExamen record in the database.
// It returns the ID of the newly inserted tipoExamen record.
func (r repository) CrearTipoExamen(ctx context.Context, tipoExamen entity.TipoExamen) (entity.TipoExamen, error) {
	err := r.db.With(ctx).Model(&tipoExamen).Insert()
	if err != nil {
		return entity.TipoExamen{}, err
	}
	return tipoExamen, nil
}

// Create saves a new TipoExamen record in the database.
// It returns the ID of the newly inserted tipoExamen record.
func (r repository) ActualizarTipoExamen(ctx context.Context, tipoExamen entity.TipoExamen) (entity.TipoExamen, error) {
	var err error
	if tipoExamen.IdTipoExamen != 0 {
		err = r.db.With(ctx).Model(&tipoExamen).Update()
	} else {
		err = r.db.With(ctx).Model(&tipoExamen).Insert()
	}
	if err != nil {
		return entity.TipoExamen{}, err
	}
	return tipoExamen, nil
}

// GetTipoExamenPorId reads the tipoExamen with the specified ID from the database.
func (r repository) GetTipoExamenPorId(ctx context.Context, idTipoExamen int) (entity.TipoExamen, error) {
	var tipoExamen entity.TipoExamen
	err := r.db.With(ctx).Select().Model(idTipoExamen, &tipoExamen)
	return tipoExamen, err
}
