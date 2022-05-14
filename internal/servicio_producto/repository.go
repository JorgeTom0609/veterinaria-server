package servicio_producto

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access servicioProductos from the data source.
type Repository interface {
	// GetServicioProductoPorId returns the servicioProducto with the specified servicioProducto ID.
	GetServicioProductoPorId(ctx context.Context, idServicioProducto int) (entity.ServicioProducto, error)
	// GetServicioProductos returns the list servicioProductos.
	GetServicioProductos(ctx context.Context) ([]entity.ServicioProducto, error)
	GetServicioProductosConDatos(ctx context.Context) ([]ServicioProductoConDatos, error)
	CrearServicioProducto(ctx context.Context, servicioProducto entity.ServicioProducto) (entity.ServicioProducto, error)
	ActualizarServicioProducto(ctx context.Context, servicioProducto entity.ServicioProducto) (entity.ServicioProducto, error)
}

// repository persists servicioProductos in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new servicioProducto repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list servicioProductos from the database.
func (r repository) GetServicioProductos(ctx context.Context) ([]entity.ServicioProducto, error) {
	var servicioProductos []entity.ServicioProducto

	err := r.db.With(ctx).
		Select().
		From().
		All(&servicioProductos)
	if err != nil {
		return servicioProductos, err
	}
	return servicioProductos, err
}

func (r repository) GetServicioProductosConDatos(ctx context.Context) ([]ServicioProductoConDatos, error) {
	var servicioProductosConDatos []ServicioProductoConDatos
	var servicioProductos []entity.ServicioProducto

	err := r.db.With(ctx).
		Select().
		From().
		All(&servicioProductos)

	if err != nil {
		return []ServicioProductoConDatos{}, err
	}

	for i := 0; i < len(servicioProductos); i++ {
		var producto entity.Producto
		idProducto := servicioProductos[i].IdProducto
		err := r.db.With(ctx).
			Select().
			Where(dbx.HashExp{"id_producto": idProducto}).
			One(&producto)
		if err != nil {
			return []ServicioProductoConDatos{}, err
		}
		servicioProductosConDatos = append(servicioProductosConDatos, ServicioProductoConDatos{
			servicioProductos[i],
			producto,
		})
	}

	return servicioProductosConDatos, err
}

// Create saves a new ServicioProducto record in the database.
// It returns the ID of the newly inserted servicioProducto record.
func (r repository) CrearServicioProducto(ctx context.Context, servicioProducto entity.ServicioProducto) (entity.ServicioProducto, error) {
	err := r.db.With(ctx).Model(&servicioProducto).Insert()
	if err != nil {
		return entity.ServicioProducto{}, err
	}
	return servicioProducto, nil
}

// Create saves a new ServicioProducto record in the database.
// It returns the ID of the newly inserted servicioProducto record.
func (r repository) ActualizarServicioProducto(ctx context.Context, servicioProducto entity.ServicioProducto) (entity.ServicioProducto, error) {
	var err error
	if servicioProducto.IdServicioProducto != 0 {
		err = r.db.With(ctx).Model(&servicioProducto).Update()
	} else {
		err = r.db.With(ctx).Model(&servicioProducto).Insert()
	}
	if err != nil {
		return entity.ServicioProducto{}, err
	}
	return servicioProducto, nil
}

// Get reads the servicioProducto with the specified ID from the database.
func (r repository) GetServicioProductoPorId(ctx context.Context, idServicioProducto int) (entity.ServicioProducto, error) {
	var servicioProducto entity.ServicioProducto
	err := r.db.With(ctx).Select().Model(idServicioProducto, &servicioProducto)
	return servicioProducto, err
}
