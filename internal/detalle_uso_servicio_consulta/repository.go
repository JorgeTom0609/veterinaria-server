package detalle_uso_servicio_consulta

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access detallesUsoServicioConsulta from the data source.
type Repository interface {
	// GetDetalleUsoServicioConsultaPorId returns the detalleUsoServicioConsulta with the specified detalleUsoServicioConsulta ID.
	GetDetalleUsoServicioConsultaPorId(ctx context.Context, idDetalleUsoServicioConsulta int) (entity.DetalleUsoServicioConsulta, error)
	// GetDetallesUsoServicioConsulta returns the list detallesUsoServicioConsulta.
	GetDetallesUsoServicioConsulta(ctx context.Context) ([]entity.DetalleUsoServicioConsulta, error)
	CrearDetalleUsoServicioConsulta(ctx context.Context, detalleUsoServicioConsulta entity.DetalleUsoServicioConsulta) (entity.DetalleUsoServicioConsulta, error)
	ActualizarDetalleUsoServicioConsulta(ctx context.Context, detalleUsoServicioConsulta entity.DetalleUsoServicioConsulta) (entity.DetalleUsoServicioConsulta, error)
}

// repository persists detallesUsoServicioConsulta in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new detalleUsoServicioConsulta repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list detallesUsoServicioConsulta from the database.
func (r repository) GetDetallesUsoServicioConsulta(ctx context.Context) ([]entity.DetalleUsoServicioConsulta, error) {
	var detallesUsoServicioConsulta []entity.DetalleUsoServicioConsulta

	err := r.db.With(ctx).
		Select().
		From().
		All(&detallesUsoServicioConsulta)
	if err != nil {
		return detallesUsoServicioConsulta, err
	}
	return detallesUsoServicioConsulta, err
}

// Create saves a new DetalleUsoServicioConsulta record in the database.
// It returns the ID of the newly inserted detalleUsoServicioConsulta record.
func (r repository) CrearDetalleUsoServicioConsulta(ctx context.Context, detalleUsoServicioConsulta entity.DetalleUsoServicioConsulta) (entity.DetalleUsoServicioConsulta, error) {
	err := r.db.With(ctx).Model(&detalleUsoServicioConsulta).Insert()
	if err != nil {
		return entity.DetalleUsoServicioConsulta{}, err
	}
	return detalleUsoServicioConsulta, nil
}

// Create saves a new DetalleUsoServicioConsulta record in the database.
// It returns the ID of the newly inserted detalleUsoServicioConsulta record.
func (r repository) ActualizarDetalleUsoServicioConsulta(ctx context.Context, detalleUsoServicioConsulta entity.DetalleUsoServicioConsulta) (entity.DetalleUsoServicioConsulta, error) {
	var err error
	if detalleUsoServicioConsulta.IdDetalleUsoServicioConsulta != 0 {
		err = r.db.With(ctx).Model(&detalleUsoServicioConsulta).Update()
	} else {
		err = r.db.With(ctx).Model(&detalleUsoServicioConsulta).Insert()
	}
	if err != nil {
		return entity.DetalleUsoServicioConsulta{}, err
	}
	return detalleUsoServicioConsulta, nil
}

// Get reads the detalleUsoServicioConsulta with the specified ID from the database.
func (r repository) GetDetalleUsoServicioConsultaPorId(ctx context.Context, idDetalleUsoServicioConsulta int) (entity.DetalleUsoServicioConsulta, error) {
	var detalleUsoServicioConsulta entity.DetalleUsoServicioConsulta
	err := r.db.With(ctx).Select().Model(idDetalleUsoServicioConsulta, &detalleUsoServicioConsulta)
	return detalleUsoServicioConsulta, err
}
