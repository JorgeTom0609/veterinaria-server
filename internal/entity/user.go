package entity

import "database/sql"

// User represents a user.
type User struct {
	IdUsuario     int          `json:"id_usuario" db:"pk,id_usuario"`
	Nombre        string       `json:"nombre" db:"nombre"`
	Apellido      string       `json:"apellido" db:"apellido"`
	NombreUsuario string       `json:"nombre_usuario" db:"nombre_usuario"`
	Clave         string       `json:"clave" db:"clave"`
	Estado        sql.NullBool `json:"estado" db:"estado"`
}

// TableName represents the table name
func (u User) TableName() string {
	return "usuarios"
}

// GetIdUsuario returns the user ID.
func (u User) GetIdUsuario() int {
	return u.IdUsuario
}

// GetNombreusuario returns the user Nombre.
func (u User) GetNombreUsuario() string {
	return u.NombreUsuario
}

// GetEstado returns the user Nombre.
func (u User) IsEstado() sql.NullBool {
	return u.Estado
}
