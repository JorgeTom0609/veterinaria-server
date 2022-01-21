package entity

import "database/sql"

// Rol represents a rol.
type Rol struct {
	IdRol       int          `json:"id_rol" db:"id_rol"`
	Descripcion string       `json:"descripcion" db:"descripcion"`
	Estado      sql.NullBool `json:"estado" db:"estado"`
}

// TableName represents the table name
func (r Rol) TableName() string {
	return "roles"
}

// GetIdRol returns the rol ID.
func (r Rol) GetIdRol() int {
	return r.IdRol
}

// GetDescripcion returns the user Nombre.
func (r Rol) GetDescripcion() string {
	return r.Descripcion
}

// IsEstado returns the user Nombre.
func (u Rol) IsEstado() sql.NullBool {
	return u.Estado
}
