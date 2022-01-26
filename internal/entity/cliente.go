package entity

type Cliente struct {
	IdCliente int     `json:"id_cliente" db:"id_cliente"`
	Nombres   string  `json:"nombres" db:"nombres"`
	Apellidos string  `json:"apellidos" db:"apellidos"`
	Correo    *string `json:"correo" db:"correo"`
	Telefono  *string `json:"telefono" db:"telefono"`
	Direccion *string `json:"direccion" db:"direccion"`
}
