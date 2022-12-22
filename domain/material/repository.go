package material

import "gorm.io/gorm"

type IRepository interface {
	Save(material Material) (Material, error)
	FecthAll() ([]Material, error)
	FecthById(id int) (Material, error)
	Update(material Material) (Material, error)
	Delete(material Material) (Material, error)
}

type repository struct {
	DB *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(material Material) (Material, error) {
	err := r.DB.Create(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}

func (r *repository) FecthAll() ([]Material, error) {
	var materials []Material

	err := r.DB.Find(&materials).Error
	if err != nil {
		return materials, err
	}

	return materials, nil
}

func (r *repository) FecthById(id int) (Material, error) {
	var material Material

	err := r.DB.Where("id = ?", id).Find(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}

func (r *repository) Update(material Material) (Material, error) {
	err := r.DB.Save(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}

func (r *repository) Delete(material Material) (Material, error) {
	err := r.DB.Delete(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}
