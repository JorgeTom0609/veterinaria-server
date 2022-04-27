package entity

type StockIndividual struct {
	IdStockIndividual int     `json:"id_stock_individual" db:"pk,id_stock_individual"`
	IdLote            int     `json:"id_lote" db:"id_lote"`
	Descripcion       string  `json:"descripcion" db:"descripcion"`
	Cantidad          float32 `json:"cantidad" db:"cantidad"`
	CantidadInicial   float32 `json:"cantidad_inicial" db:"cantidad_inicial"`
}

func (r StockIndividual) TableName() string {
	return "stock_individual"
}
