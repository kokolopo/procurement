package purchase_order

import (
	"procurement/domain/purchase_request"
	"time"
)

type PurchaseOrder struct {
	Id                 int                              `gorm:"primary_key;auto_increment;not_null"`
	PurchaseRequest_id int                              `gorm:"type:int(25);not null"`
	PurchaseRequest    purchase_request.PurchaseRequest `gorm:"foreignKey:PurchaseRequest_id;not null"`
	Quantity           int                              `gorm:"type:int(25);not null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
