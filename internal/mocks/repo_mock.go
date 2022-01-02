package mocks

import (
	"errors"
	"hex_ddd_app/internal/core/domain"
)

// simple mock that satisfies CustomerRepository interface
type repoDBMock struct {
	ErrOut bool
}

func NewDBMock(errOut bool) *repoDBMock {
	return &repoDBMock{
		ErrOut: errOut,
	}
}

func (db *repoDBMock) GetCustomer(id string) (domain.Customer, error) {
	if db.ErrOut {
		return domain.Customer{}, errors.New("Invalid customer")
	}

	return domain.Customer{FirstName: "Ada", LastName: "Lovelace", SignedWaiver: true}, nil
}

func (db *repoDBMock) SaveCustomer(customer domain.Customer) error {
	if db.ErrOut {
		return errors.New("Invalid customer")
	}

	return nil
}
