package service

import (
	"endpoints/domain"
	"endpoints/dto"
)

type DefaultCustomerServer struct {
	customer domain.ICustomerRepository
}
type ICustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, error)
	GetCustomer(id int) (*dto.CustomerResponse, error)
}

func (s DefaultCustomerServer) toDtoCustomer(c []domain.Customer) []dto.CustomerResponse {
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

func (s DefaultCustomerServer) GetAllCustomers() ([]dto.CustomerResponse, error) {
	all, err := s.customer.FindAll()
	if err != nil {
		return nil, err
	}
	return s.toDtoCustomer(all), err
}

func (s DefaultCustomerServer) GetCustomer(id int) (*dto.CustomerResponse, error) {
	customer, err := s.customer.Find(id)
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

func NewCustomerService(customer domain.ICustomerRepository, account domain.IAccountRepository) DefaultCustomerServer {
	return DefaultCustomerServer{
		customer: customer,
	}
}
