package mocks

import (
	"errors"
	"hex_ddd_app/internal/client/entities"
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

func (db *repoDBMock) GetCustomer(id uint) (entities.Customer, error) {
	if db.ErrOut {
		return entities.Customer{}, errors.New("Invalid customer")
	}

	return entities.Customer{ID: id, FirstName: "Ada", LastName: "Lovelace", SignedWaiver: true}, nil
}

func (db *repoDBMock) SaveCustomer(customer *entities.Customer) error {
	if db.ErrOut {
		return errors.New("Invalid customer")
	}

	return nil
}
