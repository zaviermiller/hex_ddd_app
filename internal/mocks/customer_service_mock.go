package mocks

import (
	"errors"

	"hex_ddd_app/internal/client/entities"
	"hex_ddd_app/internal/client/ports"
)

type customerServiceMock struct {
	customerRepo ports.CustomerRepository

	ErrOut bool
}

func NewCustomerServiceMock(errOut bool) *customerServiceMock {
	return &customerServiceMock{
		ErrOut: errOut,
	}
}

func (s *customerServiceMock) Get(id uint) (entities.Customer, error) {
	if s.ErrOut {
		return entities.Customer{}, errors.New("invalid id")
	}
	return entities.Customer{ID: 1, FirstName: "Ada", LastName: "Lovelace", SignedWaiver: true}, nil
}

func (s *customerServiceMock) Create(customer *entities.Customer) error {
	customer.ID = 1

	if s.ErrOut {
		return errors.New("invalid customer")
	}

	return nil
}
