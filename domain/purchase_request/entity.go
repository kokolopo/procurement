package purchase_request

import (
	"procurement/domain/material"
	"time"
)

type PurchaseRequest struct {
	Id           int               `gorm:"primary_key;auto_increment;not_null"`
	Material_id  int               `gorm:"type:int(25);not null"`
	Material     material.Material `gorm:"foreignKey:Material_id;not null"`
	Stock_need   int               `gorm:"type:int(25);not null"`
	Requset_date time.Time         `gorm:"not null"`
}
