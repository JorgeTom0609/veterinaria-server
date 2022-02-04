package clientes

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access clientes from the data source.
type Repository interface {
	// GetClientePorId returns the cliente with the specified cliente ID.
	GetClientePorId(ctx context.Context, idCliente int) (entity.Cliente, error)
	// GetClientes returns the list clientes.
	GetClientes(ctx context.Context) ([]entity.Cliente, error)
	CrearCliente(ctx context.Context, cliente entity.Cliente) (entity.Cliente, error)
	ActualizarCliente(ctx context.Context, cliente entity.Cliente) (entity.Cliente, error)
}

// repository persists clientes in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new cliente repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list clientes from the database.
func (r repository) GetClientes(ctx context.Context) ([]entity.Cliente, error) {
	var clientes []entity.Cliente

	err := r.db.With(ctx).
		Select().
		From().
		All(&clientes)
	if err != nil {
		return clientes, err
	}
	return clientes, err
}

// Create saves a new Cliente record in the database.
// It returns the ID of the newly inserted cliente record.
func (r repository) CrearCliente(ctx context.Context, cliente entity.Cliente) (entity.Cliente, error) {
	err := r.db.With(ctx).Model(&cliente).Insert()
	if err != nil {
		return entity.Cliente{}, err
	}
	return cliente, nil
}

// Create saves a new Cliente record in the database.
// It returns the ID of the newly inserted cliente record.
func (r repository) ActualizarCliente(ctx context.Context, cliente entity.Cliente) (entity.Cliente, error) {
	var err error
	if cliente.IdCliente != 0 {
		err = r.db.With(ctx).Model(&cliente).Update()
	} else {
		err = r.db.With(ctx).Model(&cliente).Insert()
	}
	if err != nil {
		return entity.Cliente{}, err
	}
	return cliente, nil
}

// Get reads the cliente with the specified ID from the database.
func (r repository) GetClientePorId(ctx context.Context, idCliente int) (entity.Cliente, error) {
	var cliente entity.Cliente
	err := r.db.With(ctx).Select().Model(idCliente, &cliente)
	return cliente, err
}
