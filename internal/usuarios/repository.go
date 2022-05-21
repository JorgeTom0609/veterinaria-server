package usuarios

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access users from the data source.
type Repository interface {
	// GetUserPorId returns the user with the specified user ID.
	GetUserPorId(ctx context.Context, idUser int) (entity.User, error)
	// GetUsers returns the list users.
	GetUsers(ctx context.Context) ([]entity.User, error)
	CrearUser(ctx context.Context, user entity.User) (entity.User, error)
	ActualizarUser(ctx context.Context, user entity.User) (entity.User, error)
}

// repository persists users in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new user repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list users from the database.
func (r repository) GetUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	err := r.db.With(ctx).
		Select().
		From().
		All(&users)
	if err != nil {
		return users, err
	}
	return users, err
}

// Create saves a new User record in the database.
// It returns the ID of the newly inserted user record.
func (r repository) CrearUser(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.db.With(ctx).Model(&user).Insert()
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// Create saves a new User record in the database.
// It returns the ID of the newly inserted user record.
func (r repository) ActualizarUser(ctx context.Context, user entity.User) (entity.User, error) {
	var err error
	if user.IdUsuario != 0 {
		err = r.db.With(ctx).Model(&user).Update()
	} else {
		err = r.db.With(ctx).Model(&user).Insert()
	}
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// Get reads the user with the specified ID from the database.
func (r repository) GetUserPorId(ctx context.Context, idUser int) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().Model(idUser, &user)
	return user, err
}
