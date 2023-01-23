package purchase_order

import "gorm.io/gorm"

type IRepository interface {
	Save(data PurchaseOrder) (PurchaseOrder, error)
	FecthAll() ([]PurchaseOrder, error)
	FecthById(id int) (PurchaseOrder, error)
	Delete(purchaseOrder PurchaseOrder) (PurchaseOrder, error)
	Update(purchaseOrder PurchaseOrder) (PurchaseOrder, error)
}

type repository struct {
	DB *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(data PurchaseOrder) (PurchaseOrder, error) {
	err := r.DB.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) FecthAll() ([]PurchaseOrder, error) {
	var pOrder []PurchaseOrder

	err := r.DB.Preload("PurchaseRequest").Preload("Supplier").Find(&pOrder).Error
	if err != nil {
		return pOrder, err
	}

	return pOrder, nil
}

func (r *repository) FecthById(id int) (PurchaseOrder, error) {
	var pOrder PurchaseOrder

	err := r.DB.Where("id = ?", id).Preload("PurchaseRequest").Preload("Supplier").Find(&pOrder).Error
	if err != nil {
		return pOrder, err
	}

	return pOrder, nil
}

func (r *repository) Delete(purchaseOrder PurchaseOrder) (PurchaseOrder, error) {
	err := r.DB.Delete(&purchaseOrder).Error
	if err != nil {
		return purchaseOrder, err
	}

	return purchaseOrder, nil
}

func (r *repository) Update(purchaseOrder PurchaseOrder) (PurchaseOrder, error) {
	err := r.DB.Save(&purchaseOrder).Error
	if err != nil {
		return purchaseOrder, err
	}

	return purchaseOrder, nil
}
