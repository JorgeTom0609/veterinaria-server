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

// TableName represents the table name
func (r Acceso) TableName() string {
	return "accesos"
}

// GetIdAcceso returns the acceso IdAcceso.
func (a Acceso) GetIdAcceso() int {
	return a.IdAcceso
}

// GetIdAccesoPadre returns the acceso IdAccesoPadre.
func (a Acceso) GetIdAccesoPadre() *int {
	return a.IdAccesoPadre
}

// GetDescripcion returns the acceso Nombre.
func (a Acceso) GetDescripcion() string {
	return a.Descripcion
}

// GetRuta returns the acceso Ruta.
func (a Acceso) GetRuta() string {
	return a.Ruta
}

// GetIcono returns the acceso Icono.
func (a Acceso) GetIcono() *string {
	return a.Icono
}

// IsPrincipal returns the acceso Principal.
func (a Acceso) IsPrincipal() sql.NullBool {
	return a.Principal
}
