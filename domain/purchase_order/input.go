package purchase_order

type InputPO struct {
	PurchaseRequest_id int `json:"purchase_request_id" binding:"required"`
	Quantity           int `json:"quantity" binding:"required"`
}

type InputEditPO struct {
	Quantity           int `json:"quantity" binding:"required"`
}