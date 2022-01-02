package mocks

import (
	"errors"
	"fmt"

	"hex_ddd_app/internal/core/domain"
	"hex_ddd_app/internal/core/ports"
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

func (s *customerServiceMock) Get(id string) (domain.Customer, error) {
	if s.ErrOut {
		return domain.Customer{}, errors.New("invalid id")
	}
	return domain.Customer{ID: "1", FirstName: "Ada", LastName: "Lovelace", SignedWaiver: true}, nil
}

func (s *customerServiceMock) Create(customer domain.Customer) (domain.Customer, error) {
	customer.ID = fmt.Sprintf("%d", 1)

	if s.ErrOut {
		return domain.Customer{}, errors.New("invalid customer")
	}

	return customer, nil
}
