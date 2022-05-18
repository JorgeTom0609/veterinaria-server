package documento_mascota

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access documentosMascota from the data source.
type Repository interface {
	// GetDocumentoMascotaPorId returns the documentoMascota with the specified documentoMascota ID.
	GetDocumentoMascotaPorId(ctx context.Context, idDocumentoMascota int) (entity.DocumentoMascota, error)
	// GetDocumentosMascota returns the list documentosMascota.
	GetDocumentosMascota(ctx context.Context) ([]entity.DocumentoMascota, error)
	GetDocumentoMascotaPorMascota(ctx context.Context, idMascota int) ([]entity.DocumentoMascota, error)
	CrearDocumentoMascota(ctx context.Context, documentoMascota entity.DocumentoMascota) (entity.DocumentoMascota, error)
	ActualizarDocumentoMascota(ctx context.Context, documentoMascota entity.DocumentoMascota) (entity.DocumentoMascota, error)
}

// repository persists documentosMascota in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new documentoMascota repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list documentosMascota from the database.
func (r repository) GetDocumentosMascota(ctx context.Context) ([]entity.DocumentoMascota, error) {
	var documentosMascota []entity.DocumentoMascota

	err := r.db.With(ctx).
		Select().
		From().
		All(&documentosMascota)
	if err != nil {
		return documentosMascota, err
	}
	return documentosMascota, err
}

func (r repository) GetDocumentoMascotaPorMascota(ctx context.Context, idMascota int) ([]entity.DocumentoMascota, error) {
	var documentosMascota []entity.DocumentoMascota

	err := r.db.With(ctx).
		Select().
		Where(dbx.HashExp{"id_mascota": idMascota}).
		All(&documentosMascota)
	if err != nil {
		return documentosMascota, err
	}
	return documentosMascota, err
}

// Create saves a new DocumentoMascota record in the database.
// It returns the ID of the newly inserted documentoMascota record.
func (r repository) CrearDocumentoMascota(ctx context.Context, documentoMascota entity.DocumentoMascota) (entity.DocumentoMascota, error) {
	err := r.db.With(ctx).Model(&documentoMascota).Insert()
	if err != nil {
		return entity.DocumentoMascota{}, err
	}
	return documentoMascota, nil
}

// Create saves a new DocumentoMascota record in the database.
// It returns the ID of the newly inserted documentoMascota record.
func (r repository) ActualizarDocumentoMascota(ctx context.Context, documentoMascota entity.DocumentoMascota) (entity.DocumentoMascota, error) {
	var err error
	if documentoMascota.IdDocumentoMascota != 0 {
		err = r.db.With(ctx).Model(&documentoMascota).Update()
	} else {
		err = r.db.With(ctx).Model(&documentoMascota).Insert()
	}
	if err != nil {
		return entity.DocumentoMascota{}, err
	}
	return documentoMascota, nil
}

// Get reads the documentoMascota with the specified ID from the database.
func (r repository) GetDocumentoMascotaPorId(ctx context.Context, idDocumentoMascota int) (entity.DocumentoMascota, error) {
	var documentoMascota entity.DocumentoMascota
	err := r.db.With(ctx).Select().Model(idDocumentoMascota, &documentoMascota)
	return documentoMascota, err
}
