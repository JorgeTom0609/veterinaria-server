package entity

type Proveedor struct {
	IdProveedor int     `json:"id_proveedor" db:"pk,id_proveedor"`
	Descripcion string  `json:"descripcion" db:"descripcion"`
	Celular     *string `json:"celular" db:"celular"`
	Correo      *string `json:"correo" db:"correo"`
	Ruc         *string `json:"ruc" db:"ruc"`
	Direccion   *string `json:"direccion" db:"direccion"`
}

func (p Proveedor) TableName() string {
	return "proveedor"
}
