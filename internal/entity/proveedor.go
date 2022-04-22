package entity

type Proveedor struct {
	IdProveedor int    `json:"id_proveedor" db:"pk,id_proveedor"`
	Descripcion string `json:"descripcion" db:"descripcion"`
}

func (p Proveedor) TableName() string {
	return "proveedor"
}
