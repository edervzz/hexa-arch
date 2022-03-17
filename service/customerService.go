package service

import "helloworld/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(id int) (domain.Customer, error)
}

type DefaulCustomerServer struct {
	repo domain.CustomerRepository
}

func (s DefaulCustomerServer) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaulCustomerServer) GetCustomer(id int) (domain.Customer, error) {
	return s.repo.Find(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaulCustomerServer {
	return DefaulCustomerServer{repo: repository}
}
