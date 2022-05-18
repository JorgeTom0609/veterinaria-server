package entity

import (
	"time"
)

type DocumentoMascota struct {
	IdDocumentoMascota int       `json:"id_documento_mascota" db:"pk,id_documento_mascota"`
	IdMascota          int       `json:"id_mascota" db:"id_mascota"`
	IdUsuario          int       `json:"id_usuario" db:"id_usuario"`
	Nombre             string    `json:"nombre" db:"nombre"`
	Extension          string    `json:"extension" db:"extension"`
	Ruta               string    `json:"ruta" db:"ruta"`
	Descripcion        string    `json:"descripcion" db:"descripcion"`
	Fecha              time.Time `json:"fecha" db:"fecha"`
}

func (c DocumentoMascota) TableName() string {
	return "documento_mascota"
}
