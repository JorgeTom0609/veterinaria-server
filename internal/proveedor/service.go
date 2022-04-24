package proveedor

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for proveedores.
type Service interface {
	GetProveedores(ctx context.Context) ([]Proveedor, error)
	GetProveedorPorId(ctx context.Context, idProveedor int) (Proveedor, error)
	CrearProveedor(ctx context.Context, input CreateProveedorRequest) (Proveedor, error)
	ActualizarProveedor(ctx context.Context, input UpdateProveedorRequest) (Proveedor, error)
}

// Proveedores represents the data about an proveedores.
type Proveedor struct {
	entity.Proveedor
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new proveedores service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list proveedores.
func (s service) GetProveedores(ctx context.Context) ([]Proveedor, error) {
	proveedores, err := s.repo.GetProveedores(ctx)
	if err != nil {
		return nil, err
	}
	result := []Proveedor{}
	for _, item := range proveedores {
		result = append(result, Proveedor{item})
	}
	return result, nil
}

// CreateProveedorRequest represents an proveedor creation request.
type CreateProveedorRequest struct {
	Descripcion string  `json:"descripcion"`
	Celular     *string `json:"celular"`
	Correo      *string `json:"correo"`
	Ruc         *string `json:"ruc"`
	Direccion   *string `json:"direccion"`
}

type UpdateProveedorRequest struct {
	IdProveedor int     `json:"id_proveedor"`
	Descripcion string  `json:"descripcion"`
	Celular     *string `json:"celular"`
	Correo      *string `json:"correo"`
	Ruc         *string `json:"ruc"`
	Direccion   *string `json:"direccion"`
}

// Validate validates the UpdateProveedorRequest fields.
func (m UpdateProveedorRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
	)
}

// Validate validates the CreateProveedorRequest fields.
func (m CreateProveedorRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
	)
}

// CrearProveedor creates a new proveedor.
func (s service) CrearProveedor(ctx context.Context, req CreateProveedorRequest) (Proveedor, error) {
	if err := req.Validate(); err != nil {
		return Proveedor{}, err
	}
	proveedorG, err := s.repo.CrearProveedor(ctx, entity.Proveedor{
		Descripcion: req.Descripcion,
		Celular:     req.Celular,
		Correo:      req.Correo,
		Ruc:         req.Ruc,
		Direccion:   req.Direccion,
	})
	if err != nil {
		return Proveedor{}, err
	}
	return Proveedor{proveedorG}, nil
}

// ActualizarProveedor creates a new proveedor.
func (s service) ActualizarProveedor(ctx context.Context, req UpdateProveedorRequest) (Proveedor, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Proveedor{}, err
	}
	proveedorG, err := s.repo.ActualizarProveedor(ctx, entity.Proveedor{
		IdProveedor: req.IdProveedor,
		Descripcion: req.Descripcion,
		Celular:     req.Celular,
		Correo:      req.Correo,
		Ruc:         req.Ruc,
		Direccion:   req.Direccion,
	})
	if err != nil {
		return Proveedor{}, err
	}
	return Proveedor{proveedorG}, nil
}

// GetProveedorPorId returns the proveedor with the specified the proveedor ID.
func (s service) GetProveedorPorId(ctx context.Context, idProveedor int) (Proveedor, error) {
	proveedor, err := s.repo.GetProveedorPorId(ctx, idProveedor)
	if err != nil {
		return Proveedor{}, err
	}
	return Proveedor{proveedor}, nil
}
