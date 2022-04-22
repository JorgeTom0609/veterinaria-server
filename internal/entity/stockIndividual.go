package entity

type StockIndividual struct {
	IdStockIndividual int     `json:"id_stock_individual" db:"pk,id_stock_individual"`
	IdLote            int     `json:"id_lote" db:"id_lote"`
	IdUnidad          int     `json:"id_unidad" db:"id_unidad"`
	Descripcion       string  `json:"descripcion" db:"descripcion"`
	Cantidad          float32 `json:"cantidad" db:"cantidad"`
}

func (r StockIndividual) TableName() string {
	return "stock_individual"
}
