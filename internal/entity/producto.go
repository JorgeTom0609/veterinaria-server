package entity

import "database/sql"

type Producto struct {
	IdProducto   int          `json:"id_producto" db:"pk,id_producto"`
	Descripcion  string       `json:"descripcion" db:"descripcion"`
	PrecioVenta  float32      `json:"precio_venta" db:"precio_venta"`
	Iva          sql.NullBool `json:"iva" db:"iva"`
	UsoInterno   sql.NullBool `json:"uso_interno" db:"uso_interno"`
	VentaPublico sql.NullBool `json:"venta_publico" db:"venta_publico"`
	PorMedida    sql.NullBool `json:"por_medida" db:"por_medida"`
}

func (r Producto) TableName() string {
	return "producto"
}
