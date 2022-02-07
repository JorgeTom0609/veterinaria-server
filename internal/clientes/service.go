package clientes

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Service encapsulates usecase logic for clientes.
type Service interface {
	GetClientes(ctx context.Context) ([]Cliente, error)
	GetClientePorId(ctx context.Context, idCliente int) (Cliente, error)
	CrearCliente(ctx context.Context, input CreateClienteRequest) (Cliente, error)
	ActualizarCliente(ctx context.Context, input UpdateClienteRequest) (Cliente, error)
}

// Clientes represents the data about an clientes.
type Cliente struct {
	entity.Cliente
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new clientes service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list clientes.
func (s service) GetClientes(ctx context.Context) ([]Cliente, error) {
	clientes, err := s.repo.GetClientes(ctx)
	if err != nil {
		return nil, err
	}
	result := []Cliente{}
	for _, item := range clientes {
		result = append(result, Cliente{item})
	}
	return result, nil
}

// CreateClienteRequest represents an cliente creation request.
type CreateClienteRequest struct {
	Nombres   string  `json:"nombres"`
	Apellidos string  `json:"apellidos"`
	Correo    *string `json:"correo"`
	Telefono  *string `json:"telefono"`
	Direccion *string `json:"direccion"`
}

type UpdateClienteRequest struct {
	IdCliente int     `json:"id_cliente"`
	Nombres   string  `json:"nombres"`
	Apellidos string  `json:"apellidos"`
	Correo    *string `json:"correo"`
	Telefono  *string `json:"telefono"`
	Direccion *string `json:"direccion"`
}

// Validate validates the UpdateClienteRequest fields.
func (m UpdateClienteRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Nombres, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Apellidos, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Correo, is.Email),
	)
}

// Validate validates the CreateClienteRequest fields.
func (m CreateClienteRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Nombres, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Apellidos, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Correo, is.Email),
	)
}

// CrearCliente creates a new cliente.
func (s service) CrearCliente(ctx context.Context, req CreateClienteRequest) (Cliente, error) {
	if err := req.Validate(); err != nil {
		return Cliente{}, err
	}
	clienteG, err := s.repo.CrearCliente(ctx, entity.Cliente{
		Nombres:   req.Nombres,
		Apellidos: req.Apellidos,
		Correo:    req.Correo,
		Telefono:  req.Telefono,
		Direccion: req.Direccion,
	})
	if err != nil {
		return Cliente{}, err
	}
	return Cliente{clienteG}, nil
}

// ActualizarCliente creates a new cliente.
func (s service) ActualizarCliente(ctx context.Context, req UpdateClienteRequest) (Cliente, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Cliente{}, err
	}
	clienteG, err := s.repo.ActualizarCliente(ctx, entity.Cliente{
		IdCliente: req.IdCliente,
		Nombres:   req.Nombres,
		Apellidos: req.Apellidos,
		Correo:    req.Correo,
		Telefono:  req.Telefono,
		Direccion: req.Direccion,
	})
	if err != nil {
		return Cliente{}, err
	}
	return Cliente{clienteG}, nil
}

// GetClientePorId returns the cliente with the specified the cliente ID.
func (s service) GetClientePorId(ctx context.Context, idCliente int) (Cliente, error) {
	cliente, err := s.repo.GetClientePorId(ctx, idCliente)
	if err != nil {
		return Cliente{}, err
	}
	return Cliente{cliente}, nil
}
