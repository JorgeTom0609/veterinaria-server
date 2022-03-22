package entity

type ProductoVP struct {
	IdProductoVP int     `json:"id_producto_vp" db:"pk,id_producto_vp"`
	Descripcion  string  `json:"descripcion" db:"descripcion"`
	PrecioCompra float32 `json:"precio_compra" db:"precio_compra"`
	PrecioVenta  float32 `json:"precio_venta" db:"precio_venta"`
	Stock        int     `json:"stock" db:"stock"`
	StockMinimo  *int    `json:"stock_minimo" db:"stock_minimo"`
}

func (p ProductoVP) TableName() string {
	return "productos_vp"
}
