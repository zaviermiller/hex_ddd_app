package repositories

import (
	"hex_ddd_app/internal/core/domain"
	"testing"
)

func Test_NewInMemoryDB(t *testing.T) {
	db := NewInMemoryDB()
	if len(db.data) > 0 {
		t.Error("Expected NewInMemoryDB to create db with no entries")
	}
}

func Test_MemDB_GetCustomer(t *testing.T) {
	db := NewInMemoryDB()

	got, err := db.GetCustomer("1")
	if err == nil {
		t.Errorf("Expected memDB.GetCustomer to fail, instead got %v", got)
	}

	db.data["1"] = domain.Customer{ID: "1", FirstName: "Linus", LastName: "Torvalds"}

	cust, err := db.GetCustomer("1")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if cust.ID != "1" {
		t.Errorf("Returned customer primary key did not match mem db key (expected: 1, got: %s)", cust.ID)
	}
}

func Test_MemDB_SaveCustomer(t *testing.T) {
	db := NewInMemoryDB()

	err := db.SaveCustomer(domain.Customer{})

	if err == nil || err.Error() != "missing primary key" {
		t.Errorf("Invalid SaveCustomer did not fail or gave wrong error: %s", err.Error())
	}

	newCust := domain.Customer{ID: "1", FirstName: "Linus", LastName: "Torvalds"}

	err = db.SaveCustomer(newCust)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	err = db.SaveCustomer(newCust)
	if err == nil {
		t.Errorf("Expected error when trying to enter duplicate record")
	}
}
