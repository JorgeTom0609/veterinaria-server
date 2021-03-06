package detalle_servicio_consulta

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access detallesServicioConsulta from the data source.
type Repository interface {
	// GetDetalleServicioConsultaPorId returns the detalleServicioConsulta with the specified detalleServicioConsulta ID.
	GetDetalleServicioConsultaPorId(ctx context.Context, idDetalleServicioConsulta int) (entity.DetalleServicioConsulta, error)
	GetDetalleServicioConsultaPorConsulta(ctx context.Context, idConsulta int) ([]DetalleServicioConsultaConDatos, error)
	// GetDetallesServicioConsulta returns the list detallesServicioConsulta.
	GetDetallesServicioConsulta(ctx context.Context) ([]entity.DetalleServicioConsulta, error)
	CrearDetalleServicioConsulta(ctx context.Context, detalleServicioConsulta entity.DetalleServicioConsulta) (entity.DetalleServicioConsulta, error)
	ActualizarDetalleServicioConsulta(ctx context.Context, detalleServicioConsulta entity.DetalleServicioConsulta) (entity.DetalleServicioConsulta, error)
}

// repository persists detallesServicioConsulta in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new detalleServicioConsulta repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list detallesServicioConsulta from the database.
func (r repository) GetDetallesServicioConsulta(ctx context.Context) ([]entity.DetalleServicioConsulta, error) {
	var detallesServicioConsulta []entity.DetalleServicioConsulta

	err := r.db.With(ctx).
		Select().
		From().
		All(&detallesServicioConsulta)
	if err != nil {
		return detallesServicioConsulta, err
	}
	return detallesServicioConsulta, err
}

// Create saves a new DetalleServicioConsulta record in the database.
// It returns the ID of the newly inserted detalleServicioConsulta record.
func (r repository) CrearDetalleServicioConsulta(ctx context.Context, detalleServicioConsulta entity.DetalleServicioConsulta) (entity.DetalleServicioConsulta, error) {
	err := r.db.With(ctx).Model(&detalleServicioConsulta).Insert()
	if err != nil {
		return entity.DetalleServicioConsulta{}, err
	}
	return detalleServicioConsulta, nil
}

// Create saves a new DetalleServicioConsulta record in the database.
// It returns the ID of the newly inserted detalleServicioConsulta record.
func (r repository) ActualizarDetalleServicioConsulta(ctx context.Context, detalleServicioConsulta entity.DetalleServicioConsulta) (entity.DetalleServicioConsulta, error) {
	var err error
	if detalleServicioConsulta.IdDetalleServicioConsulta != 0 {
		err = r.db.With(ctx).Model(&detalleServicioConsulta).Update()
	} else {
		err = r.db.With(ctx).Model(&detalleServicioConsulta).Insert()
	}
	if err != nil {
		return entity.DetalleServicioConsulta{}, err
	}
	return detalleServicioConsulta, nil
}

// Get reads the detalleServicioConsulta with the specified ID from the database.
func (r repository) GetDetalleServicioConsultaPorId(ctx context.Context, idDetalleServicioConsulta int) (entity.DetalleServicioConsulta, error) {
	var detalleServicioConsulta entity.DetalleServicioConsulta
	err := r.db.With(ctx).Select().Model(idDetalleServicioConsulta, &detalleServicioConsulta)
	return detalleServicioConsulta, err
}

func (r repository) GetDetalleServicioConsultaPorConsulta(ctx context.Context, idConsulta int) ([]DetalleServicioConsultaConDatos, error) {
	var detallesServicioConsulta []entity.DetalleServicioConsulta
	var detallesServicioConsultaConDatos []DetalleServicioConsultaConDatos = []DetalleServicioConsultaConDatos{}
	err := r.db.With(ctx).
		Select().
		From().
		Where(dbx.HashExp{"id_consulta": idConsulta}).
		All(&detallesServicioConsulta)
	if err != nil {
		return []DetalleServicioConsultaConDatos{}, err
	}
	for i := 0; i < len(detallesServicioConsulta); i++ {
		var servicio string = ""
		err := r.db.With(ctx).
			Select("descripcion").
			From("servicios").
			Where(dbx.HashExp{"id_servicio": detallesServicioConsulta[i].IdServicio}).
			Row(&servicio)
		if err != nil {
			return []DetalleServicioConsultaConDatos{}, err
		}
		detallesServicioConsultaConDatos = append(detallesServicioConsultaConDatos, DetalleServicioConsultaConDatos{
			DetalleServicioConsulta: detallesServicioConsulta[i],
			Servicio:                servicio,
		})
	}
	return detallesServicioConsultaConDatos, err
}
