package servicios

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access servicios from the data source.
type Repository interface {
	// GetServicioPorId returns the servicio with the specified servicio ID.
	GetServicioPorId(ctx context.Context, idServicio int) (entity.Servicio, error)
	// GetServicios returns the list servicios.
	GetServicios(ctx context.Context) ([]entity.Servicio, error)
	GetServicioPorEspecie(ctx context.Context, idEspecie int, modo int) ([]ServicioTieneProductos, error)
	GetServiciosConProductos(ctx context.Context) ([]ServicioTieneProductos, error)
	CrearServicio(ctx context.Context, servicio entity.Servicio) (entity.Servicio, error)
	ActualizarServicio(ctx context.Context, servicio entity.Servicio) (entity.Servicio, error)
}

// repository persists servicios in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new servicio repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list servicios from the database.
func (r repository) GetServicios(ctx context.Context) ([]entity.Servicio, error) {
	var servicios []entity.Servicio

	err := r.db.With(ctx).
		Select().
		From().
		All(&servicios)
	if err != nil {
		return servicios, err
	}
	return servicios, err
}

func (r repository) GetServicioPorEspecie(ctx context.Context, idEspecie int, modo int) ([]ServicioTieneProductos, error) {
	var servicios []entity.Servicio
	var serviciosTieneProductos []ServicioTieneProductos

	var err error
	if modo == 1 {
		err = r.db.With(ctx).
			Select().
			Where(dbx.HashExp{"id_especie": idEspecie}).
			AndWhere(dbx.NewExp("aplica_consulta = true")).
			All(&servicios)
	} else if modo == 2 {
		err = r.db.With(ctx).
			Select().
			Where(dbx.HashExp{"id_especie": idEspecie}).
			AndWhere(dbx.NewExp("aplica_hospitalizacion = true")).
			All(&servicios)
	} else {
		err = r.db.With(ctx).
			Select().
			Where(dbx.HashExp{"id_especie": idEspecie}).
			All(&servicios)
	}

	if err != nil {
		return []ServicioTieneProductos{}, err
	}

	for i := 0; i < len(servicios); i++ {
		idServicio := servicios[i].IdServicio
		cant := 0
		err := r.db.With(ctx).
			Select("count(id_servicio)").
			From("servicio_producto").
			Where(dbx.HashExp{"id_servicio": idServicio}).
			AndWhere(dbx.HashExp{"estado": "A"}).
			Row(&cant)
		if err != nil {
			return []ServicioTieneProductos{}, err
		}
		serviciosTieneProductos = append(serviciosTieneProductos, ServicioTieneProductos{
			Servicio:         servicios[i],
			CantidadProducto: cant,
		})
	}

	return serviciosTieneProductos, err
}

func (r repository) GetServiciosConProductos(ctx context.Context) ([]ServicioTieneProductos, error) {
	var servicios []entity.Servicio
	var serviciosTieneProductos []ServicioTieneProductos

	err := r.db.With(ctx).
		Select().
		From().
		All(&servicios)
	if err != nil {
		return []ServicioTieneProductos{}, err
	}

	for i := 0; i < len(servicios); i++ {
		idServicio := servicios[i].IdServicio
		cant := 0
		err := r.db.With(ctx).
			Select("count(id_servicio)").
			From("servicio_producto").
			Where(dbx.HashExp{"id_servicio": idServicio}).
			AndWhere(dbx.HashExp{"estado": "A"}).
			Row(&cant)
		if err != nil {
			return []ServicioTieneProductos{}, err
		}
		serviciosTieneProductos = append(serviciosTieneProductos, ServicioTieneProductos{
			Servicio:         servicios[i],
			CantidadProducto: cant,
		})
	}

	return serviciosTieneProductos, err
}

// Create saves a new Servicio record in the database.
// It returns the ID of the newly inserted servicio record.
func (r repository) CrearServicio(ctx context.Context, servicio entity.Servicio) (entity.Servicio, error) {
	err := r.db.With(ctx).Model(&servicio).Insert()
	if err != nil {
		return entity.Servicio{}, err
	}
	return servicio, nil
}

// Create saves a new Servicio record in the database.
// It returns the ID of the newly inserted servicio record.
func (r repository) ActualizarServicio(ctx context.Context, servicio entity.Servicio) (entity.Servicio, error) {
	var err error
	if servicio.IdServicio != 0 {
		err = r.db.With(ctx).Model(&servicio).Update()
	} else {
		err = r.db.With(ctx).Model(&servicio).Insert()
	}
	if err != nil {
		return entity.Servicio{}, err
	}
	return servicio, nil
}

// Get reads the servicio with the specified ID from the database.
func (r repository) GetServicioPorId(ctx context.Context, idServicio int) (entity.Servicio, error) {
	var servicio entity.Servicio
	err := r.db.With(ctx).Select().Model(idServicio, &servicio)
	return servicio, err
}
