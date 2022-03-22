package entity

type DetalleFactura struct {
	IdDetalleFactura int     `json:"id_detalle_factura" db:"pk,id_detalle_factura"`
	IdFactura        int     `json:"id_factura" db:"id_factura"`
	IdProductoVp     int     `json:"id_producto_vp" db:"id_producto_vp"`
	Cantidad         int     `json:"cantidad" db:"cantidad"`
	Valor            float32 `json:"valor" db:"valor"`
}

func (d DetalleFactura) TableName() string {
	return "detalles_factura"
}
