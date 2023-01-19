package purchase_request

import (
	"time"
)

// merupakan list permintaan dari gudang barang mentah
// gudang BM meminta barang ke procurement untuk

type PurchaseRequest struct {
	Id           int    `gorm:"primary_key;auto_increment;not_null"`
	MaterialName string `gorm:"type:varchar(255);not null"`
	Stock_need   int    `gorm:"type:int(25);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
