package auth

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"veterinaria-server/internal/entity"
	"veterinaria-server/internal/errors"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	"github.com/dgrijalva/jwt-go"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"golang.org/x/crypto/bcrypt"
)

// Service encapsulates the authentication logic.
type Service interface {
	// authenticate authenticates a user using username and password.
	// It returns a JWT token if authentication succeeds. Otherwise, an error is returned.
	Login(ctx context.Context, username, password string) (string, error)
}

// Identity represents an authenticated user identity.
type Identity interface {
	// GetIdUsuario returns the user ID.
	GetIdUsuario() int
	// GetNombreUsuario returns the user Username.
	GetNombreUsuario() string
	// IsEstado returns the user status
	IsEstado() sql.NullBool
	GetNombres() string
}

type service struct {
	db              *dbcontext.DB
	signingKey      string
	tokenExpiration int
	logger          log.Logger
}

// NewService creates a new authentication service.
func NewService(db *dbcontext.DB, signingKey string, tokenExpiration int, logger log.Logger) Service {
	return service{db, signingKey, tokenExpiration, logger}
}

// Login authenticates a user and generates a JWT token if authentication succeeds.
// Otherwise, an error is returned.
func (s service) Login(ctx context.Context, username, password string) (string, error) {
	if identity := s.authenticate(ctx, username, password); identity != nil {
		return s.generateJWT(identity)
	}
	return "", errors.Unauthorized("")
}

// authenticate authenticates a user using username and password.
// If username and password are correct, an identity is returned. Otherwise, nil is returned.
func (s service) authenticate(ctx context.Context, username, password string) Identity {
	logger := s.logger.With(ctx, "user", username)

	user := entity.User{}

	if err := s.db.With(ctx).Select().Where(dbx.HashExp{"nombre_usuario": username, "estado": true}).One(&user); err != nil {
		fmt.Println(err)
		return nil
	}

	/* para encriptar la pass
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Clave), bcrypt.MinCost)
	fmt.Println(string(hash))
	*/

	if err := bcrypt.CompareHashAndPassword([]byte(user.Clave), []byte(password)); err != nil {
		fmt.Println(err)
		logger.Infof("authentication failed")
		return nil
	}
	logger.Infof("authentication successful")
	u := entity.User{IdUsuario: user.IdUsuario, NombreUsuario: user.NombreUsuario, Estado: user.Estado, Nombre: user.Nombre, Apellido: user.Apellido}
	return u
}

// generateJWT generates a JWT that encodes an identity.
func (s service) generateJWT(identity Identity) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       identity.GetIdUsuario(),
		"username": identity.GetNombreUsuario(),
		"exp":      time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
		"nombres":  identity.GetNombres(),
	}).SignedString([]byte(s.signingKey))
}
