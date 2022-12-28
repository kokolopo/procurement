package purchase_request

import "gorm.io/gorm"

type IRepository interface {
	Save(data PurchaseRequest) (PurchaseRequest, error)
	FecthAll() ([]PurchaseRequest, error)
	FecthById(id int) (PurchaseRequest, error)
	// Update(material Material) (Material, error)
	// Delete(material Material) (Material, error)
}

type repository struct {
	DB *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(data PurchaseRequest) (PurchaseRequest, error) {
	err := r.DB.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) FecthAll() ([]PurchaseRequest, error) {
	var pRequests []PurchaseRequest

	err := r.DB.Find(&pRequests).Error
	if err != nil {
		return pRequests, err
	}

	return pRequests, nil
}

func (r *repository) FecthById(id int) (PurchaseRequest, error) {
	var pRequest PurchaseRequest

	err := r.DB.Where("id = ?", id).Find(&pRequest).Error
	if err != nil {
		return pRequest, err
	}

	return pRequest, nil
}