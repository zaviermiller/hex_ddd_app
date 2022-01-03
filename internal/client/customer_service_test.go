package client

import (
	"hex_ddd_app/internal/client/entities"
	"hex_ddd_app/internal/mocks"
	"testing"
)

func Test_NewCustomerService(t *testing.T) {
	mockDb := mocks.NewDBMock(false)
	srv := NewCustomerService(mockDb)

	if srv.customerRepo != mockDb {
		t.Errorf("expected new CustomerService.repository to be %v, got: %v", mockDb, srv.customerRepo)

	}
}

func Test_CustomerService_Create(t *testing.T) {
	dbMock := mocks.NewDBMock(true)
	srv := NewCustomerService(dbMock)

	cust := &entities.Customer{FirstName: "Alan", LastName: "Turing"}
	err := srv.Create(cust)
	if err == nil {
		t.Error("Expected CustomerService.Create to err")
	}

	dbMock.ErrOut = false
	err = srv.Create(cust)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

}

func Test_CustomerService_Get(t *testing.T) {
	dbMock := mocks.NewDBMock(true)
	srv := NewCustomerService(dbMock)

	_, err := srv.Get(1)
	if err == nil {
		t.Error("Expected CustomerService.Get to err")
	}

	dbMock.ErrOut = false
	cust, _ := srv.Get(1)

	if c, _ := dbMock.GetCustomer(1); c != cust {
		t.Errorf("Unexpected customer returned from CustomerService.Get, expected: %v, got: %v", c, cust)
	}
}
