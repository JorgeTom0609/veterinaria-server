package rol

import (
	"context"
	"database/sql"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for roles.
type Service interface {
	GetRoles(ctx context.Context) ([]Rol, error)
	GetRolPorId(ctx context.Context, idRol int) (Rol, error)
	CrearRol(ctx context.Context, input CreateRolRequest) (Rol, error)
	ActualizarRol(ctx context.Context, input UpdateRolRequest) (Rol, error)
}

// Roles represents the data about an roles.
type Rol struct {
	entity.Rol
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new roles service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list roles.
func (s service) GetRoles(ctx context.Context) ([]Rol, error) {
	roles, err := s.repo.GetRoles(ctx)
	if err != nil {
		return nil, err
	}
	result := []Rol{}
	for _, item := range roles {
		result = append(result, Rol{item})
	}
	return result, nil
}

// CreateRolRequest represents an rol creation request.
type CreateRolRequest struct {
	Descripcion string       `json:"descripcion"`
	Estado      sql.NullBool `json:"estado"`
}

type UpdateRolRequest struct {
	IdRol       int          `json:"id_rol"`
	Descripcion string       `json:"descripcion"`
	Estado      sql.NullBool `json:"estado"`
}

// Validate validates the UpdateRolRequest fields.
func (m UpdateRolRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 1000)),
	)
}

// Validate validates the CreateRolRequest fields.
func (m CreateRolRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 1000)),
	)
}

// CrearRol creates a new rol.
func (s service) CrearRol(ctx context.Context, req CreateRolRequest) (Rol, error) {
	if err := req.Validate(); err != nil {
		return Rol{}, err
	}
	rolG, err := s.repo.CrearRol(ctx, entity.Rol{
		Descripcion: req.Descripcion,
		Estado:      req.Estado,
	})
	if err != nil {
		return Rol{}, err
	}
	return Rol{rolG}, nil
}

// ActualizarRol creates a new rol.
func (s service) ActualizarRol(ctx context.Context, req UpdateRolRequest) (Rol, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Rol{}, err
	}
	rolG, err := s.repo.ActualizarRol(ctx, entity.Rol{
		IdRol:       req.IdRol,
		Descripcion: req.Descripcion,
		Estado:      req.Estado,
	})
	if err != nil {
		return Rol{}, err
	}
	return Rol{rolG}, nil
}

// GetRolPorId returns the rol with the specified the rol ID.
func (s service) GetRolPorId(ctx context.Context, idRol int) (Rol, error) {
	rol, err := s.repo.GetRolPorId(ctx, idRol)
	if err != nil {
		return Rol{}, err
	}
	return Rol{rol}, nil
}
