package entity

type DetalleFactura struct {
	IdDetalleFactura int     `json:"id_detalle_factura" db:"pk,id_detalle_factura"`
	IdFactura        int     `json:"id_factura" db:"id_factura"`
	IdReferencia     int     `json:"id_referencia" db:"id_referencia"`
	Tabla            string  `json:"tabla" db:"tabla"`
	Cantidad         float32 `json:"cantidad" db:"cantidad"`
	Valor            float32 `json:"valor" db:"valor"`
}

func (d DetalleFactura) TableName() string {
	return "detalles_factura"
}
