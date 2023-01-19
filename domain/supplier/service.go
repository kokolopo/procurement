package supplier

type IService interface {
	Save(input AddSuplier) (Supplier, error)
	GetAll() ([]Supplier, error)
}

type service struct {
	repository IRepository
}

func NewSupplierService(repository IRepository) *service {
	return &service{repository}
}

func (s *service) Save(input AddSuplier) (Supplier, error) {
	var data Supplier

	//tangkap nilai dari inputan
	data.Email = input.Email
	data.SupplierName = input.SupplierName
	data.Address = input.Address
	data.Phone = input.Phone

	//save data yang sudah dimapping kedalam struct PR
	supplier, err := s.repository.Save(data)
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (s *service) GetAll() ([]Supplier, error) {
	materials, err := s.repository.FecthAll()
	if err != nil {
		return nil, err
	}

	return materials, nil
}