package repositories

import (
	"errors"
	"hex_ddd_app/internal/client/entities"
)

type memDB struct {
	data      map[uint]interface{}
	currentID int
}

func NewInMemoryDB() *memDB {
	dbMap := make(map[uint]interface{})
	return &memDB{
		data:      dbMap,
		currentID: 1,
	}
}

func (db *memDB) GetCustomer(id uint) (entities.Customer, error) {
	custIface := db.data[id]
	cust, ok := custIface.(entities.Customer)
	if !ok {
		return entities.Customer{}, errors.New("Invalid id given")
	}
	return cust, nil
}

func (db *memDB) SaveCustomer(customer *entities.Customer) error {
	customer.ID = uint(db.currentID)

	// TODO: should update. or create crud actions
	if _, ok := db.data[customer.ID]; ok {
		customer.ID = 0
		return errors.New("primary key already in database")
	}
	db.data[customer.ID] = *customer

	db.currentID += 1

	return nil
}
