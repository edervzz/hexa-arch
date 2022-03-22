package service

import (
	"endpoints/domain"
	"endpoints/dto"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, error)
	GetCustomer(id int) (*dto.CustomerResponse, error)
}

type DefaulCustomerServer struct {
	repo domain.CustomerRepository
}

func (s DefaulCustomerServer) toDto(c []domain.Customer) []dto.CustomerResponse {
	var cr dto.CustomerResponse
	var r []dto.CustomerResponse
	for _, v := range c {
		cr.ID = v.ID_customer
		cr.Fullname = v.Name
		cr.City = v.City
		cr.Zipcode = v.Zipcode
		r = append(r, cr)
	}
	return r
}

func (s DefaulCustomerServer) GetAllCustomers() ([]dto.CustomerResponse, error) {
	all, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return s.toDto(all), err
}

func (s DefaulCustomerServer) GetCustomer(id int) (*dto.CustomerResponse, error) {
	customer, err := s.repo.Find(id)
	if err != nil {
		return nil, err
	}
	var r *dto.CustomerResponse = &dto.CustomerResponse{
		ID:       customer.ID_customer,
		Fullname: customer.Name,
		City:     customer.City,
		Zipcode:  customer.Zipcode,
	}
	return r, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaulCustomerServer {
	return DefaulCustomerServer{repo: repository}
}
