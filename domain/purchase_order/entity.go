package purchase_order

import (
	"procurement/domain/material"
	"time"
)

type PurchaseOrder struct {
	Id          int               `gorm:"primary_key;auto_increment;not_null"`
	Material_id int               `gorm:"type:int(25);not null"`
	Material    material.Material `gorm:"foreignKey:Material_id;not null"`
	Quantity    int               `gorm:"type:int(25);not null"`
	Price       int               `gorm:"type:int(25);not null"`
	Total       int               `gorm:"type:int(25);not null"`
	PO_date     time.Time         `gorm:"not null"`
}
