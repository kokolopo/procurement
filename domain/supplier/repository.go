package supplier

import "gorm.io/gorm"

type IRepository interface {
	Save(supplier Supplier) (Supplier, error)
	FecthAll() ([]Supplier, error)
	FecthById(id int) (Supplier, error)
	Update(supplier Supplier) (Supplier, error)
	Delete(supplier Supplier) (Supplier, error)
}

type repository struct {
	DB *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(supplier Supplier) (Supplier, error) {
	err := r.DB.Create(&supplier).Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r *repository) FecthAll() ([]Supplier, error) {
	var materials []Supplier

	err := r.DB.Find(&materials).Error
	if err != nil {
		return materials, err
	}

	return materials, nil
}

func (r *repository) FecthById(id int) (Supplier, error) {
	var material Supplier

	err := r.DB.Where("id = ?", id).Find(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}

func (r *repository) Update(material Supplier) (Supplier, error) {
	err := r.DB.Save(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}

func (r *repository) Delete(material Supplier) (Supplier, error) {
	err := r.DB.Delete(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}
