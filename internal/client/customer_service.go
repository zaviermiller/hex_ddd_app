package client

import (
	"errors"

	"hex_ddd_app/internal/client/entities"
	"hex_ddd_app/internal/client/ports"
)

// services provide needed functionality for all driver adapters
type customerService struct {
	customerRepo ports.CustomerRepository
}

func NewCustomerService(repo ports.CustomerRepository) *customerService {
	return &customerService{
		customerRepo: repo,
	}
}

func (s *customerService) Get(id uint) (entities.Customer, error) {
	cust, err := s.customerRepo.GetCustomer(id)
	if err != nil {
		return entities.Customer{}, errors.New("Unable to retrieve customer with given ID")
	}

	return cust, nil
}

func (s *customerService) Create(customer *entities.Customer) error {
	err := s.customerRepo.SaveCustomer(customer)
	if err != nil {
		return err
	}

	return nil
}
