package entity

type Mascota struct {
	IdMascota int `json:"id_mascota" db:"id_mascota"`
	IdEspecie int `json:"id_especie" db:"id_especie"`
	IdCliente int `json:"id_cliente" db:"id_cliente"`
	IdGenero  int `json:"id_genero" db:"id_genero"`
}
