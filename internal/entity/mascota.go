package entity

type Mascota struct {
	IdMascota int     `json:"id_mascota" db:"pk,id_mascota"`
	IdEspecie int     `json:"id_especie" db:"id_especie"`
	IdCliente int     `json:"id_cliente" db:"id_cliente"`
	IdGenero  int     `json:"id_genero" db:"id_genero"`
	Nombre    *string `json:"nombre" db:"nombre"`
	Raza      *string `json:"raza" db:"raza"`
	Color     *string `json:"color" db:"color"`
}

func (m Mascota) TableName() string {
	return "mascotas"
}
