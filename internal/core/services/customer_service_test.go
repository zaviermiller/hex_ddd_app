package services

import (
	"fmt"
	"hex_ddd_app/internal/core/domain"
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

	cust := domain.Customer{FirstName: "Alan", LastName: "Turing"}
	_, err := srv.Create(cust)
	if err == nil {
		t.Error("Expected CustomerService.Create to err")
	}

	dbMock.ErrOut = false
	cust, _ = srv.Create(cust)

	if cust.ID != "1" {
		t.Errorf("Expected newly created CustomerService to set first customer id to 1. got %s", cust.ID)
	}

	if cust.ID == fmt.Sprintf("%d", srv.currentId) {
		t.Errorf("CustomerService.currentId should not be the same as newly created customer id")
	}
}

func Test_CustomerService_Get(t *testing.T) {
	dbMock := mocks.NewDBMock(true)
	srv := NewCustomerService(dbMock)

	_, err := srv.Get("1")
	if err == nil {
		t.Error("Expected CustomerService.Get to err")
	}

	dbMock.ErrOut = false
	cust, _ := srv.Get("1")

	if c, _ := dbMock.GetCustomer("1"); c != cust {
		t.Errorf("Unexpected customer returned from CustomerService.Get, expected: %v, got: %v", c, cust)
	}
}
