package entity

// Rol represents a rol.
type UsuarioRol struct {
	IdUsuarioRol int `json:"id_usuario_rol" db:"pk,id_usuario_rol"`
	IdRol        int `json:"id_rol" db:"id_rol"`
	IdUsuario    int `json:"id_usuario" db:"id_usuario"`
}

func (ur UsuarioRol) TableName() string {
	return "usuario_rol"
}
