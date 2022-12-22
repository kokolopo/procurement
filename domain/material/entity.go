package material

import "procurement/domain/supplier"

type Material struct {
	Id            int               `gorm:"primary_key;auto_increment;not_null"`
	Supplier_id   int               `gorm:"type:int(25);not null"`
	Supplier      supplier.Supplier `gorm:"foreignKey:Supplier_id;not null"`
	Material_name string            `gorm:"type:varchar(100);not null"`
	Material_type string            `gorm:"type:varchar(100);not null"`
	Stock         int               `gorm:"type:int(100);not null"`
}
