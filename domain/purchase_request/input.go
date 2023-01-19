package purchase_request

type RequestForm struct {
	MaterialName string `json:"material_name" binding:"required"`
	StockNeed    int    `json:"stock_need" binding:"required"`
}
