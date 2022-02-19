package examen_mascota

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access examenesMascota from the data source.
type Repository interface {
	// GetExamenMascotaPorId returns the examenesMascota with the specified examenesMascota ID.
	GetExamenMascotaPorId(ctx context.Context, idExamenMascota int) (entity.ExamenMascota, error)
	// GetExamenesMascota returns the list examenesMascota.
	GetExamenesMascota(ctx context.Context) ([]entity.ExamenMascota, error)
	GetExamenesMascotaPorMascotayEstado(ctx context.Context, idExamenMascota int, estado string) ([]entity.ExamenMascota, error)
	CrearExamenMascota(ctx context.Context, examenesMascota entity.ExamenMascota) (entity.ExamenMascota, error)
	ActualizarExamenMascota(ctx context.Context, examenesMascota entity.ExamenMascota) (entity.ExamenMascota, error)
}

// repository persists examenesMascota in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new examenesMascota repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list examenesMascota from the database.
func (r repository) GetExamenesMascota(ctx context.Context) ([]entity.ExamenMascota, error) {
	var examenesMascota []entity.ExamenMascota

	err := r.db.With(ctx).
		Select().
		From().
		All(&examenesMascota)
	if err != nil {
		return examenesMascota, err
	}
	return examenesMascota, err
}

func (r repository) GetExamenesMascotaPorMascotayEstado(ctx context.Context, idMascota int, estado string) ([]entity.ExamenMascota, error) {
	var examenesMascota []entity.ExamenMascota

	err := r.db.With(ctx).
		Select().
		From().
		Where(dbx.HashExp{"id_Mascota": idMascota}).
		AndWhere(dbx.HashExp{"estado": estado}).
		All(&examenesMascota)
	if err != nil {
		return examenesMascota, err
	}
	return examenesMascota, err
}

// Create saves a new ExamenMascota record in the database.
// It returns the ID of the newly inserted examenesMascota record.
func (r repository) CrearExamenMascota(ctx context.Context, examenesMascota entity.ExamenMascota) (entity.ExamenMascota, error) {
	err := r.db.With(ctx).Model(&examenesMascota).Insert()
	if err != nil {
		return entity.ExamenMascota{}, err
	}
	return examenesMascota, nil
}

// Create saves a new ExamenMascota record in the database.
// It returns the ID of the newly inserted examenesMascota record.
func (r repository) ActualizarExamenMascota(ctx context.Context, examenesMascota entity.ExamenMascota) (entity.ExamenMascota, error) {
	var err error
	if examenesMascota.IdExamenMascota != 0 {
		err = r.db.With(ctx).Model(&examenesMascota).Update()
	} else {
		err = r.db.With(ctx).Model(&examenesMascota).Insert()
	}
	if err != nil {
		return entity.ExamenMascota{}, err
	}
	return examenesMascota, nil
}

// Get reads the examenesMascota with the specified ID from the database.
func (r repository) GetExamenMascotaPorId(ctx context.Context, idExamenMascota int) (entity.ExamenMascota, error) {
	var examenesMascota entity.ExamenMascota
	err := r.db.With(ctx).Select().Model(idExamenMascota, &examenesMascota)
	return examenesMascota, err
}
