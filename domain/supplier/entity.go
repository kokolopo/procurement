package supplier

type Supplier struct {
	Id            int    `gorm:"primary_key;auto_increment;not_null"`
	Email         string `gorm:"uniqueIndex;not null"`
	Username      string `gorm:"type:varchar(255);not null"`
	Password      string `gorm:"type:longtext;not null"`
	Supplier_name string `gorm:"type:varchar(255);not null"`
	Address       string `gorm:"type:longtext;not null"`
	City          string `gorm:"type:varchar(255);not null"`
	Phone         string `gorm:"type:varchar(255);not null"`
}
