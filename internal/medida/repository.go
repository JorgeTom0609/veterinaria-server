package medida

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access medidas from the data source.
type Repository interface {
	// GetMedidaPorId returns the medida with the specified medida ID.
	GetMedidaPorId(ctx context.Context, idMedida int) (entity.Medida, error)
	// GetMedidas returns the list medidas.
	GetMedidas(ctx context.Context) ([]entity.Medida, error)
	CrearMedida(ctx context.Context, medida entity.Medida) (entity.Medida, error)
	ActualizarMedida(ctx context.Context, medida entity.Medida) (entity.Medida, error)
}

// repository persists medidas in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new medida repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list medidas from the database.
func (r repository) GetMedidas(ctx context.Context) ([]entity.Medida, error) {
	var medidas []entity.Medida

	err := r.db.With(ctx).
		Select().
		All(&medidas)
	if err != nil {
		return medidas, err
	}
	return medidas, err
}

// Create saves a new Medida record in the database.
// It returns the ID of the newly inserted medida record.
func (r repository) CrearMedida(ctx context.Context, medida entity.Medida) (entity.Medida, error) {
	err := r.db.With(ctx).Model(&medida).Insert()
	if err != nil {
		return entity.Medida{}, err
	}
	return medida, nil
}

// Create saves a new Medida record in the database.
// It returns the ID of the newly inserted medida record.
func (r repository) ActualizarMedida(ctx context.Context, medida entity.Medida) (entity.Medida, error) {
	var err error
	if medida.IdMedida != 0 {
		err = r.db.With(ctx).Model(&medida).Update()
	} else {
		err = r.db.With(ctx).Model(&medida).Insert()
	}
	if err != nil {
		return entity.Medida{}, err
	}
	return medida, nil
}

// Get reads the medida with the specified ID from the database.
func (r repository) GetMedidaPorId(ctx context.Context, idMedida int) (entity.Medida, error) {
	var medida entity.Medida
	err := r.db.With(ctx).Select().Model(idMedida, &medida)
	return medida, err
}
