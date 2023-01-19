package purchase_request

import "procurement/domain/material"

type IService interface {
	GetAll() ([]PurchaseRequest, error)
	GetById(id int) (PurchaseRequest, error)
	Save(input RequestForm) (PurchaseRequest, error)
}

type service struct {
	repository   IRepository
	materialRepo material.IRepository
}

func NewPRService(repository IRepository, materialRepo material.IRepository) *service {
	return &service{repository, materialRepo}
}

func (s *service) Save(input RequestForm) (PurchaseRequest, error) {
	var data PurchaseRequest

	//tangkap nilai dari inputan
	data.MaterialName = input.MaterialName
	data.Stock_need = input.StockNeed

	//save data yang sudah dimapping kedalam struct PR
	newPR, err := s.repository.Save(data)
	if err != nil {
		return newPR, err
	}

	return newPR, nil
}

// GetAll implements IService
func (s *service) GetAll() ([]PurchaseRequest, error) {
	requests, err := s.repository.FecthAll()
	if err != nil {
		return nil, err
	}

	return requests, nil
}

// GetById implements IService
func (s *service) GetById(id int) (PurchaseRequest, error) {
	request, err := s.repository.FecthById(id)
	if err != nil {
		return request, err
	}

	return request, nil
}
