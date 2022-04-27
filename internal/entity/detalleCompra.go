package entity

type DetalleCompra struct {
	IdDetalleCompra int     `json:"id_detalle_compra" db:"pk,id_detalle_compra"`
	IdCompra        int     `json:"id_compra" db:"id_compra"`
	IdLote          int     `json:"id_lote" db:"id_lote"`
	Cantidad        int     `json:"cantidad" db:"cantidad"`
	Valor           float32 `json:"valor" db:"valor"`
}

func (d DetalleCompra) TableName() string {
	return "detalles_compra"
}
