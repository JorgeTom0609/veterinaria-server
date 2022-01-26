package entity

import "database/sql"

// Rol represents a rol.
type Rol struct {
	IdRol       int          `json:"id_rol" db:"pk,id_rol"`
	Descripcion string       `json:"descripcion" db:"descripcion"`
	Estado      sql.NullBool `json:"estado" db:"estado"`
}

func (r Rol) TableName() string {
	return "roles"
}
