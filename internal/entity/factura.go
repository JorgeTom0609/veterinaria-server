package entity

import "time"

type Factura struct {
	IdFactura int       `json:"id_factura" db:"pk,id_factura"`
	IdCliente int       `json:"id_cliente" db:"id_cliente"`
	Fecha     time.Time `json:"fecha" db:"fecha"`
	Valor     float32   `json:"valor" db:"valor"`
}

func (f Factura) TableName() string {
	return "facturas"
}
