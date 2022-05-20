package hospitalizacion

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access hospitalizaciones from the data source.
type Repository interface {
	// GetHospitalizacionPorId returns the hospitalizacion with the specified hospitalizacion ID.
	GetHospitalizacionPorId(ctx context.Context, idHospitalizacion int) (entity.Hospitalizacion, error)
	// GetHospitalizaciones returns the list hospitalizaciones.
	GetHospitalizaciones(ctx context.Context) ([]entity.Hospitalizacion, error)
	GetHospitalizacionesActivas(ctx context.Context) ([]HospitalizacionesActivas, error)
	CrearHospitalizacion(ctx context.Context, hospitalizacion entity.Hospitalizacion) (entity.Hospitalizacion, error)
	ActualizarHospitalizacion(ctx context.Context, hospitalizacion entity.Hospitalizacion) (entity.Hospitalizacion, error)
}

// repository persists hospitalizaciones in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new hospitalizacion repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list hospitalizaciones from the database.
func (r repository) GetHospitalizaciones(ctx context.Context) ([]entity.Hospitalizacion, error) {
	var hospitalizaciones []entity.Hospitalizacion

	err := r.db.With(ctx).
		Select().
		From().
		All(&hospitalizaciones)
	if err != nil {
		return hospitalizaciones, err
	}
	return hospitalizaciones, err
}

func (r repository) GetHospitalizacionesActivas(ctx context.Context) ([]HospitalizacionesActivas, error) {
	var hospitalizaciones []entity.Hospitalizacion
	var hospitalizacionesActivas []HospitalizacionesActivas = []HospitalizacionesActivas{}

	err := r.db.With(ctx).
		Select().
		Where(dbx.NewExp("estado_hospitalizacion = 'ACTIVA'")).
		All(&hospitalizaciones)
	if err != nil {
		return []HospitalizacionesActivas{}, err
	}

	for i := 0; i < len(hospitalizaciones); i++ {
		var mascota entity.Mascota
		var especie entity.Especie
		var consulta entity.Consulta

		err := r.db.With(ctx).
			Select().
			From("consulta").
			Where(dbx.HashExp{"id_consulta": hospitalizaciones[i].IdConsulta}).
			One(&consulta)
		if err != nil {
			return []HospitalizacionesActivas{}, err
		}

		err = r.db.With(ctx).
			Select().
			From("mascotas").
			Where(dbx.HashExp{"id_mascota": consulta.IdMascota}).
			One(&mascota)
		if err != nil {
			return []HospitalizacionesActivas{}, err
		}

		err = r.db.With(ctx).
			Select().
			From("mascotas as m").
			InnerJoin("especies as e", dbx.NewExp("e.id_especie = m.id_especie")).
			Where(dbx.HashExp{"m.id_mascota": mascota.IdMascota}).
			One(&especie)
		if err != nil {
			return []HospitalizacionesActivas{}, err
		}

		hospitalizacionesActivas = append(hospitalizacionesActivas, HospitalizacionesActivas{Hospitalizacion: hospitalizaciones[i], Mascota: mascota, Especie: especie, Consulta: consulta})
	}

	return hospitalizacionesActivas, err
}

// Create saves a new Hospitalizacion record in the database.
// It returns the ID of the newly inserted hospitalizacion record.
func (r repository) CrearHospitalizacion(ctx context.Context, hospitalizacion entity.Hospitalizacion) (entity.Hospitalizacion, error) {
	err := r.db.With(ctx).Model(&hospitalizacion).Insert()
	if err != nil {
		return entity.Hospitalizacion{}, err
	}
	return hospitalizacion, nil
}

// Create saves a new Hospitalizacion record in the database.
// It returns the ID of the newly inserted hospitalizacion record.
func (r repository) ActualizarHospitalizacion(ctx context.Context, hospitalizacion entity.Hospitalizacion) (entity.Hospitalizacion, error) {
	var err error
	if hospitalizacion.IdHospitalizacion != 0 {
		err = r.db.With(ctx).Model(&hospitalizacion).Update()
	} else {
		err = r.db.With(ctx).Model(&hospitalizacion).Insert()
	}
	if err != nil {
		return entity.Hospitalizacion{}, err
	}
	return hospitalizacion, nil
}

// Get reads the hospitalizacion with the specified ID from the database.
func (r repository) GetHospitalizacionPorId(ctx context.Context, idHospitalizacion int) (entity.Hospitalizacion, error) {
	var hospitalizacion entity.Hospitalizacion
	err := r.db.With(ctx).Select().Model(idHospitalizacion, &hospitalizacion)
	return hospitalizacion, err
}
