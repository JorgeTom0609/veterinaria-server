package entity

type Cliente struct {
	IdCliente    int     `json:"id_cliente" db:"pk,id_cliente"`
	Nombres      string  `json:"nombres" db:"nombres"`
	Apellidos    string  `json:"apellidos" db:"apellidos"`
	Cedula       string  `json:"cedula" db:"cedula"`
	Correo       *string `json:"correo" db:"correo"`
	Telefono     *string `json:"telefono" db:"telefono"`
	Direccion    *string `json:"direccion" db:"direccion"`
	Nacionalidad *string `json:"nacionalidad" db:"nacionalidad"`
}

func (c Cliente) TableName() string {
	return "clientes"
}
