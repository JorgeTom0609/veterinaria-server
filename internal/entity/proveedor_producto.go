package entity

type ProveedorProducto struct {
	IdProveedorProducto int     `json:"id_proveedor_producto" db:"pk,id_proveedor_producto"`
	IdProveedor         int     `json:"id_proveedor" db:"id_proveedor"`
	IdProducto          int     `json:"id_producto" db:"id_producto"`
	PrecioCompra        float32 `json:"precio_compra" db:"precio_compra"`
}

func (p ProveedorProducto) TableName() string {
	return "proveedor_producto"
}
