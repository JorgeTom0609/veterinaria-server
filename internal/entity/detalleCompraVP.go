package entity

type DetalleCompraVP struct {
	IdDetalleCompraVP int     `json:"id_detalle_compra_vp" db:"pk,id_detalle_compra_vp"`
	IdCompra          int     `json:"id_compra" db:"id_compra"`
	IdProductoVp      int     `json:"id_producto_vp" db:"id_producto_vp"`
	Cantidad          int     `json:"cantidad" db:"cantidad"`
	Valor             float32 `json:"valor" db:"valor"`
}

func (d DetalleCompraVP) TableName() string {
	return "detalles_factura"
}
