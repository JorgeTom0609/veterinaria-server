package usuarios

import (
	"context"
	"database/sql"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for users.
type Service interface {
	GetUsers(ctx context.Context) ([]User, error)
	GetUserPorId(ctx context.Context, idUser int) (User, error)
	CrearUser(ctx context.Context, input CreateUserRequest) (User, error)
	ActualizarUser(ctx context.Context, input UpdateUserRequest) (User, error)
}

// Users represents the data about an users.
type User struct {
	entity.User
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new users service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list users.
func (s service) GetUsers(ctx context.Context) ([]User, error) {
	users, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	result := []User{}
	for _, item := range users {
		result = append(result, User{item})
	}
	return result, nil
}

// CreateUserRequest represents an user creation request.
type CreateUserRequest struct {
	Nombre        string       `json:"nombre"`
	Apellido      string       `json:"apellido"`
	NombreUsuario string       `json:"nombre_usuario"`
	Clave         string       `json:"clave"`
	Estado        sql.NullBool `json:"estado"`
}

type UpdateUserRequest struct {
	IdUsuario     int          `json:"id_usuario"`
	Nombre        string       `json:"nombre"`
	Apellido      string       `json:"apellido"`
	NombreUsuario string       `json:"nombre_usuario"`
	Clave         string       `json:"clave"`
	Estado        sql.NullBool `json:"estado"`
}

// Validate validates the UpdateUserRequest fields.
func (m UpdateUserRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Nombre, validation.Required, validation.Length(0, 1000)),
		validation.Field(&m.Apellido, validation.Required, validation.Length(0, 1000)),
		validation.Field(&m.NombreUsuario, validation.Required, validation.Length(0, 1000)),
		validation.Field(&m.Clave, validation.Required, validation.Length(0, 1000)),
	)
}

// Validate validates the CreateUserRequest fields.
func (m CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Nombre, validation.Required, validation.Length(0, 1000)),
		validation.Field(&m.Apellido, validation.Required, validation.Length(0, 1000)),
		validation.Field(&m.NombreUsuario, validation.Required, validation.Length(0, 1000)),
		validation.Field(&m.Clave, validation.Required, validation.Length(0, 1000)),
	)
}

// CrearUser creates a new user.
func (s service) CrearUser(ctx context.Context, req CreateUserRequest) (User, error) {
	if err := req.Validate(); err != nil {
		return User{}, err
	}
	userG, err := s.repo.CrearUser(ctx, entity.User{
		Nombre:        req.Nombre,
		Apellido:      req.Apellido,
		NombreUsuario: req.NombreUsuario,
		Clave:         req.Clave,
		Estado:        req.Estado,
	})
	if err != nil {
		return User{}, err
	}
	return User{userG}, nil
}

// ActualizarUser creates a new user.
func (s service) ActualizarUser(ctx context.Context, req UpdateUserRequest) (User, error) {
	if err := req.ValidateUpdate(); err != nil {
		return User{}, err
	}
	userG, err := s.repo.ActualizarUser(ctx, entity.User{
		IdUsuario:     req.IdUsuario,
		Nombre:        req.Nombre,
		Apellido:      req.Apellido,
		NombreUsuario: req.NombreUsuario,
		Clave:         req.Clave,
		Estado:        req.Estado,
	})
	if err != nil {
		return User{}, err
	}
	return User{userG}, nil
}

// GetUserPorId returns the user with the specified the user ID.
func (s service) GetUserPorId(ctx context.Context, idUser int) (User, error) {
	user, err := s.repo.GetUserPorId(ctx, idUser)
	if err != nil {
		return User{}, err
	}
	return User{user}, nil
}
