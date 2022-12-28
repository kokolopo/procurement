package purchase_request

type RequestForm struct {
	MaterialId  int    `json:"material_id" binding:"required"`
	StockNeed   int    `json:"stock_need" binding:"required"`
}
