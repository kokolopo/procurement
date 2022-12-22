package supplier

type IService interface {
}

type service struct {
	repository IRepository
}

func NewSupplierService(repository IRepository) *service {
	return &service{repository}
}
