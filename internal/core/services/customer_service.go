package services

import (
	"errors"
	"fmt"

	"hex_ddd_app/internal/core/domain"
	"hex_ddd_app/internal/core/ports"
)

type customerService struct {
	customerRepo ports.CustomerRepository

	currentId int
}

func NewCustomerService(repo ports.CustomerRepository) *customerService {
	return &customerService{
		customerRepo: repo,
		currentId:    1,
	}
}

func (s *customerService) Get(id string) (domain.Customer, error) {
	cust, err := s.customerRepo.GetCustomer(id)
	if err != nil {
		return domain.Customer{}, errors.New("Unable to retrieve customer with given ID")
	}

	return cust, nil
}

func (s *customerService) Create(customer domain.Customer) (domain.Customer, error) {
	customer.ID = fmt.Sprintf("%d", s.currentId)

	err := s.customerRepo.SaveCustomer(customer)
	if err != nil {
		return domain.Customer{}, err
	}

	s.currentId += 1

	return customer, nil
}
