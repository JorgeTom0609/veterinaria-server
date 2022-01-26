package entity

import "database/sql"

// Rol represents a rol.
type Acceso struct {
	IdAcceso      int          `json:"id_acceso" db:"id_acceso"`
	IdAccesoPadre *int         `json:"id_acceso_padre" db:"id_acceso_padre"`
	Descripcion   string       `json:"descripcion" db:"descripcion"`
	Ruta          string       `json:"ruta" db:"ruta"`
	Icono         *string      `json:"icono" db:"icono"`
	Principal     sql.NullBool `json:"principal" db:"principal"`
}
