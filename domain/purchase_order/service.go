package purchase_order

type IService interface {
	GetAll() ([]PurchaseOrder, error)
	GetById(id int) (PurchaseOrder, error)
	Save(input InputPO) (PurchaseOrder, error)
	Delete(id int) (PurchaseOrder, error)
	Update(id int, input InputEditPO) (PurchaseOrder, error)
}

type service struct {
	repository IRepository
}

func NewPOService(repository IRepository) *service {
	return &service{repository}
}

func (s *service) Save(input InputPO) (PurchaseOrder, error) {
	var data PurchaseOrder

	//tangkap nilai dari inputan
	data.PurchaseRequest_id = input.PurchaseRequest_id
	data.Quantity = input.Quantity

	//save data yang sudah dimapping kedalam struct PR
	newPR, err := s.repository.Save(data)
	if err != nil {
		return newPR, err
	}

	return newPR, nil
}

// GetAll implements IService
func (s *service) GetAll() ([]PurchaseOrder, error) {
	requests, err := s.repository.FecthAll()
	if err != nil {
		return nil, err
	}

	return requests, nil
}

// GetById implements IService
func (s *service) GetById(id int) (PurchaseOrder, error) {
	request, err := s.repository.FecthById(id)
	if err != nil {
		return request, err
	}

	return request, nil
}

func (s *service) Delete(id int) (PurchaseOrder, error) {
	pO, err := s.repository.FecthById(id)
	if err != nil {
		return pO, err
	}

	data, err := s.repository.Delete(pO)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *service) Update(id int, input InputEditPO) (PurchaseOrder, error) {
	data, errItem := s.repository.FecthById(id)
	if errItem != nil {
		return data, errItem
	}

	data.Quantity = input.Quantity

	updatedItem, errUpdate := s.repository.Update(data)
	if errUpdate != nil {
		return updatedItem, errUpdate
	}

	return updatedItem, nil
}
