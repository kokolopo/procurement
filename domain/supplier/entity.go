package supplier

import "time"

type Supplier struct {
	Id           int    `gorm:"primary_key;auto_increment;not_null"`
	Email        string `gorm:"uniqueIndex(255);not null"`
	SupplierName string `gorm:"type:varchar(255);not null"`
	Address      string `gorm:"type:longtext;not null"`
	Phone        string `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
