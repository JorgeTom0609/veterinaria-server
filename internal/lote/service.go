package lote

import (
	"context"
	"time"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for lotes.
type Service interface {
	GetLotes(ctx context.Context) ([]Lote, error)
	GetLotePorId(ctx context.Context, idLote int) (Lote, error)
	CrearLote(ctx context.Context, input CreateLoteRequest) (Lote, error)
	ActualizarLote(ctx context.Context, input UpdateLoteRequest) (Lote, error)
}

// Lotes represents the data about an lotes.
type Lote struct {
	entity.Lote
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new lotes service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list lotes.
func (s service) GetLotes(ctx context.Context) ([]Lote, error) {
	lotes, err := s.repo.GetLotes(ctx)
	if err != nil {
		return nil, err
	}
	result := []Lote{}
	for _, item := range lotes {
		result = append(result, Lote{item})
	}
	return result, nil
}

// CreateLoteRequest represents an lote creation request.
type CreateLoteRequest struct {
	IdProveedorProducto int        `json:"id_proveedor_producto"`
	FechaCaducidad      *time.Time `json:"fecha_caducidad"`
	Stock               int        `json:"stock"`
	Descripcion         string     `json:"descripcion"`
}

type UpdateLoteRequest struct {
	IdLote              int        `json:"id_lote"`
	IdProveedorProducto int        `json:"id_proveedor_producto"`
	FechaCaducidad      *time.Time `json:"fecha_caducidad"`
	Stock               int        `json:"stock"`
	Descripcion         string     `json:"descripcion"`
}

// Validate validates the UpdateLoteRequest fields.
func (m UpdateLoteRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.IdProveedorProducto, validation.Required),
	)
}

// Validate validates the CreateLoteRequest fields.
func (m CreateLoteRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.IdProveedorProducto, validation.Required),
	)
}

// CrearLote creates a new lote.
func (s service) CrearLote(ctx context.Context, req CreateLoteRequest) (Lote, error) {
	if err := req.Validate(); err != nil {
		return Lote{}, err
	}
	loteG, err := s.repo.CrearLote(ctx, entity.Lote{
		IdProveedorProducto: req.IdProveedorProducto,
		FechaCaducidad:      req.FechaCaducidad,
		Stock:               req.Stock,
		Descripcion:         req.Descripcion,
	})
	if err != nil {
		return Lote{}, err
	}
	return Lote{loteG}, nil
}

// ActualizarLote creates a new lote.
func (s service) ActualizarLote(ctx context.Context, req UpdateLoteRequest) (Lote, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Lote{}, err
	}
	loteG, err := s.repo.ActualizarLote(ctx, entity.Lote{
		IdLote:              req.IdLote,
		IdProveedorProducto: req.IdProveedorProducto,
		FechaCaducidad:      req.FechaCaducidad,
		Stock:               req.Stock,
		Descripcion:         req.Descripcion,
	})
	if err != nil {
		return Lote{}, err
	}
	return Lote{loteG}, nil
}

// GetLotePorId returns the lote with the specified the lote ID.
func (s service) GetLotePorId(ctx context.Context, idLote int) (Lote, error) {
	lote, err := s.repo.GetLotePorId(ctx, idLote)
	if err != nil {
		return Lote{}, err
	}
	return Lote{lote}, nil
}
