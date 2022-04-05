package consultas

import (
	"context"
	"database/sql"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access consultas from the data source.
type Repository interface {
	// GetConsultaPorId returns the consulta with the specified consulta ID.
	GetConsultaPorId(ctx context.Context, idConsulta int) (entity.Consulta, error)
	GetConsultaActiva(ctx context.Context, idUsuario int) (entity.Consulta, error)
	// GetConsultas returns the list consultas.
	GetConsultas(ctx context.Context) ([]entity.Consulta, error)
	CrearConsulta(ctx context.Context, consulta entity.Consulta) (entity.Consulta, error)
	ActualizarConsulta(ctx context.Context, consulta entity.Consulta) (entity.Consulta, error)
}

// repository persists consultas in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new consulta repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list consultas from the database.
func (r repository) GetConsultas(ctx context.Context) ([]entity.Consulta, error) {
	var consultas []entity.Consulta

	err := r.db.With(ctx).
		Select().
		From().
		OrderBy("apellidos asc").
		All(&consultas)
	if err != nil {
		return consultas, err
	}
	return consultas, err
}

// Create saves a new Consulta record in the database.
// It returns the ID of the newly inserted consulta record.
func (r repository) CrearConsulta(ctx context.Context, consulta entity.Consulta) (entity.Consulta, error) {
	err := r.db.With(ctx).Model(&consulta).Insert()
	if err != nil {
		return entity.Consulta{}, err
	}
	return consulta, nil
}

// Create saves a new Consulta record in the database.
// It returns the ID of the newly inserted consulta record.
func (r repository) ActualizarConsulta(ctx context.Context, consulta entity.Consulta) (entity.Consulta, error) {
	var err error
	if consulta.IdConsulta != 0 {
		err = r.db.With(ctx).Model(&consulta).Update()
	} else {
		err = r.db.With(ctx).Model(&consulta).Insert()
	}
	if err != nil {
		return entity.Consulta{}, err
	}
	return consulta, nil
}

// Get reads the consulta with the specified ID from the database.
func (r repository) GetConsultaPorId(ctx context.Context, idConsulta int) (entity.Consulta, error) {
	var consulta entity.Consulta
	err := r.db.With(ctx).Select().Model(idConsulta, &consulta)
	return consulta, err
}

func (r repository) GetConsultaActiva(ctx context.Context, idUsuario int) (entity.Consulta, error) {
	var consulta entity.Consulta
	err := r.db.With(ctx).Select().Where(dbx.HashExp{"id_usuario": idUsuario}).
		AndWhere(dbx.HashExp{"estado_consulta": "ACTIVA"}).One(&consulta)
	if err == sql.ErrNoRows {
		err = nil
	}
	return consulta, err
}
