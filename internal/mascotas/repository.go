package mascotas

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access mascotas from the data source.
type Repository interface {
	// GetMascotasPorCliente returns the list mascotas by Cliente.
	GetMascotasPorCliente(ctx context.Context, idCliente int) ([]entity.Mascota, error)
	// GetMascotaPorId returns the mascota with the specified mascota ID.
	GetMascotaPorId(ctx context.Context, idMascota int) (entity.Mascota, error)
	CrearMascota(ctx context.Context, mascota entity.Mascota) (entity.Mascota, error)
	ActualizarMascotaPorGrupo(ctx context.Context, mascota entity.Mascota) (entity.Mascota, error)
}

// repository persists mascotas in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new mascota repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// GetMascotasPorCliente returns the list mascotas by Cliente.
func (r repository) GetMascotasPorCliente(ctx context.Context, idCliente int) ([]entity.Mascota, error) {
	var mascotas []entity.Mascota
	err := r.db.With(ctx).
		Select().
		Where(dbx.HashExp{"id_cliente": idCliente}).
		All(&mascotas)
	if err != nil {
		return mascotas, err
	}
	return mascotas, err
}

// Create saves a new Mascota record in the database.
// It returns the ID of the newly inserted mascota record.
func (r repository) CrearMascota(ctx context.Context, mascota entity.Mascota) (entity.Mascota, error) {
	err := r.db.With(ctx).Model(&mascota).Insert()
	if err != nil {
		return entity.Mascota{}, err
	}
	return mascota, nil
}

// Create saves a new Mascota record in the database.
// It returns the ID of the newly inserted mascota record.
func (r repository) ActualizarMascotaPorGrupo(ctx context.Context, mascota entity.Mascota) (entity.Mascota, error) {
	var err error
	if mascota.IdMascota != 0 {
		err = r.db.With(ctx).Model(&mascota).Update()
	} else {
		err = r.db.With(ctx).Model(&mascota).Insert()
	}
	if err != nil {
		return entity.Mascota{}, err
	}
	return mascota, nil
}

// GetMascotaPorId reads the mascota with the specified ID from the database.
func (r repository) GetMascotaPorId(ctx context.Context, idMascota int) (entity.Mascota, error) {
	var mascota entity.Mascota
	err := r.db.With(ctx).Select().Model(idMascota, &mascota)
	return mascota, err
}
