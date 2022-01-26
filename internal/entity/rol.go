package entity

import "database/sql"

// Rol represents a rol.
type Rol struct {
	IdRol       int          `json:"id_rol" db:"id_rol"`
	Descripcion string       `json:"descripcion" db:"descripcion"`
	Estado      sql.NullBool `json:"estado" db:"estado"`
}
