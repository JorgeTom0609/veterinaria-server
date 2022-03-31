package compra

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access compras from the data source.
type Repository interface {
	// GetCompraPorId returns the compra with the specified compra ID.
	GetCompraPorId(ctx context.Context, idCompra int) (entity.Compras, error)
	// GetCompras returns the list compras.
	GetCompras(ctx context.Context) ([]entity.Compras, error)
	GetComprasConDatos(ctx context.Context) ([]ComprasConDatos, error)
	CrearCompra(ctx context.Context, compra entity.Compras) (entity.Compras, error)
	ActualizarCompra(ctx context.Context, compra entity.Compras) (entity.Compras, error)
}

// repository persists compras in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new compra repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list compras from the database.
func (r repository) GetCompras(ctx context.Context) ([]entity.Compras, error) {
	var compras []entity.Compras

	err := r.db.With(ctx).
		Select().
		From().
		All(&compras)
	if err != nil {
		return compras, err
	}
	return compras, err
}

func (r repository) GetComprasConDatos(ctx context.Context) ([]ComprasConDatos, error) {
	var compras []entity.Compras
	var comprasConDatos []ComprasConDatos = []ComprasConDatos{}
	var compradorNombre, compradorApellido string

	err := r.db.With(ctx).
		Select().
		All(&compras)

	for i := 0; i < len(compras); i++ {
		idUsuario := compras[i].IdUsuario
		err := r.db.With(ctx).
			Select("nombre", "apellido").
			From("usuarios").
			Where(dbx.HashExp{"id_usuario": idUsuario}).
			Row(&compradorNombre, &compradorApellido)
		if err != nil {
			return []ComprasConDatos{}, err
		}

		comprasConDatos = append(comprasConDatos, ComprasConDatos{
			compras[i],
			compradorApellido + " " + compradorNombre,
		})
	}

	if err != nil {
		return []ComprasConDatos{}, err
	}
	return comprasConDatos, err
}

// Create saves a new Compras record in the database.
// It returns the ID of the newly inserted compra record.
func (r repository) CrearCompra(ctx context.Context, compra entity.Compras) (entity.Compras, error) {
	err := r.db.With(ctx).Model(&compra).Insert()
	if err != nil {
		return entity.Compras{}, err
	}
	return compra, nil
}

// Create saves a new Compras record in the database.
// It returns the ID of the newly inserted compra record.
func (r repository) ActualizarCompra(ctx context.Context, compra entity.Compras) (entity.Compras, error) {
	var err error
	if compra.IdCompra != 0 {
		err = r.db.With(ctx).Model(&compra).Update()
	} else {
		err = r.db.With(ctx).Model(&compra).Insert()
	}
	if err != nil {
		return entity.Compras{}, err
	}
	return compra, nil
}

// Get reads the compra with the specified ID from the database.
func (r repository) GetCompraPorId(ctx context.Context, idCompra int) (entity.Compras, error) {
	var compra entity.Compras
	err := r.db.With(ctx).Select().Model(idCompra, &compra)
	return compra, err
}
