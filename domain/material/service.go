package material

type IService interface {
	GetAll() ([]Material, error)
	GetById(id int) (Material, error)
}

type service struct {
	repository IRepository
}

func NewMaterialService(repository IRepository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Material, error) {
	materials, err := s.repository.FecthAll()
	if err != nil {
		return nil, err
	}

	return materials, nil
}

func (s *service) GetById(id int) (Material, error) {
	material, err := s.repository.FecthById(id)
	if err != nil {
		return material, err
	}

	return material, nil
}
