package compra

import (
	"context"
	"time"
	"veterinaria-server/internal/detalle_compra_vp"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for compras.
type Service interface {
	GetCompras(ctx context.Context) ([]Compras, error)
	GetComprasConDatos(ctx context.Context) ([]ComprasConDatos, error)
	GetCompraPorId(ctx context.Context, idCompra int) (Compras, error)
	CrearCompra(ctx context.Context, input CreateCompraRequest) (Compras, error)
	ActualizarCompra(ctx context.Context, input UpdateCompraRequest) (Compras, error)
}

// Compras represents the data about an compras.
type Compras struct {
	entity.Compras
}

type ComprasConDatos struct {
	entity.Compras
	Comprador string `json:"comprador"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new compras service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list compras.
func (s service) GetCompras(ctx context.Context) ([]Compras, error) {
	compras, err := s.repo.GetCompras(ctx)
	if err != nil {
		return nil, err
	}
	result := []Compras{}
	for _, item := range compras {
		result = append(result, Compras{item})
	}
	return result, nil
}

func (s service) GetComprasConDatos(ctx context.Context) ([]ComprasConDatos, error) {
	compras, err := s.repo.GetComprasConDatos(ctx)
	if err != nil {
		return nil, err
	}
	return compras, nil
}

// CreateCompraRequest represents an compra creation request.
type CreateCompraRequest struct {
	IdUsuario   int       `json:"id_usuario"`
	Fecha       time.Time `json:"fecha"`
	Valor       float32   `json:"valor"`
	Descripcion *string   `json:"descripcion"`
}

type UpdateCompraRequest struct {
	IdUsuario   int       `json:"id_usuario"`
	IdCompra    int       `json:"id_compra"`
	Fecha       time.Time `json:"fecha"`
	Valor       float32   `json:"valor"`
	Descripcion *string   `json:"descripcion"`
}

type CreateCompraConDetallesRequest struct {
	Compra            CreateCompraRequest                              `json:"compra"`
	DetallesComprasVP []detalle_compra_vp.CreateDetalleCompraVPRequest `json:"detalles_compra_vp"`
}

// Validate validates the UpdateCompraRequest fields.
func (m UpdateCompraRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdUsuario, validation.Required))
}

// Validate validates the CreateCompraRequest fields.
func (m CreateCompraRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdUsuario, validation.Required))
}

// CrearCompra creates a new compra.
func (s service) CrearCompra(ctx context.Context, req CreateCompraRequest) (Compras, error) {
	if err := req.Validate(); err != nil {
		return Compras{}, err
	}
	compraG, err := s.repo.CrearCompra(ctx, entity.Compras{
		IdUsuario:   req.IdUsuario,
		Fecha:       req.Fecha,
		Valor:       req.Valor,
		Descripcion: req.Descripcion,
	})
	if err != nil {
		return Compras{}, err
	}
	return Compras{compraG}, nil
}

// ActualizarCompra creates a new compra.
func (s service) ActualizarCompra(ctx context.Context, req UpdateCompraRequest) (Compras, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Compras{}, err
	}
	compraG, err := s.repo.ActualizarCompra(ctx, entity.Compras{
		IdCompra:    req.IdCompra,
		IdUsuario:   req.IdUsuario,
		Fecha:       req.Fecha,
		Valor:       req.Valor,
		Descripcion: req.Descripcion,
	})
	if err != nil {
		return Compras{}, err
	}
	return Compras{compraG}, nil
}

// GetCompraPorId returns the compra with the specified the compra ID.
func (s service) GetCompraPorId(ctx context.Context, idCompra int) (Compras, error) {
	compra, err := s.repo.GetCompraPorId(ctx, idCompra)
	if err != nil {
		return Compras{}, err
	}
	return Compras{compra}, nil
}
