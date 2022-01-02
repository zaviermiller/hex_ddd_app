package repositories

import (
	"errors"

	"hex_ddd_app/internal/core/domain"
)

type memDB struct {
	data map[string]interface{}
}

func NewInMemoryDB() *memDB {
	dbMap := make(map[string]interface{})
	return &memDB{
		data: dbMap,
	}
}

func (db *memDB) GetCustomer(id string) (domain.Customer, error) {
	custIface := db.data[id]
	cust, ok := custIface.(domain.Customer)
	if !ok {
		return domain.Customer{}, errors.New("Invalid id given")
	}
	return cust, nil
}

func (db *memDB) SaveCustomer(customer domain.Customer) error {
	if _, ok := db.data[customer.ID]; ok {
		return errors.New("primary key already in database")
	}
	db.data[customer.ID] = customer

	return nil
}
