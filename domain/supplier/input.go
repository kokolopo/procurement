package supplier

type AddSuplier struct {
	Email        string `json:"email" binding:"required"`
	SupplierName string `json:"supplier_name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
}
