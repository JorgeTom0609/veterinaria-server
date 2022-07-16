package entity

import (
	"time"
)

type Lote struct {
	IdLote              int        `json:"id_lote" db:"pk,id_lote"`
	IdProveedorProducto int        `json:"id_proveedor_producto" db:"id_proveedor_producto"`
	FechaCaducidad      *time.Time `json:"fecha_caducidad" db:"fecha_caducidad"`
	Stock               int        `json:"stock" db:"stock"`
	Descripcion         string     `json:"descripcion" db:"descripcion"`
	CodigoBarra         *string    `json:"codigo_barra" db:"codigo_barra"`
}

func (r Lote) TableName() string {
	return "lote"
}
