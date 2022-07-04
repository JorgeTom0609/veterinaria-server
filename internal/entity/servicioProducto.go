package entity

type ServicioProducto struct {
	IdServicioProducto int      `json:"id_servicio_producto" db:"pk,id_servicio_producto"`
	IdServicio         int      `json:"id_servicio" db:"id_servicio"`
	IdProducto         int      `json:"id_producto" db:"id_producto"`
	Cantidad           float32  `json:"cantidad" db:"cantidad"`
	Razon              *float32 `json:"razon" db:"razon"`
	Estado             string   `json:"estado" db:"estado"`
}

func (c ServicioProducto) TableName() string {
	return "servicio_producto"
}
