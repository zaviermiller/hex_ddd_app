package repositories

import (
	"hex_ddd_app/db"
	"hex_ddd_app/internal/client/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pgDb struct {
	db *gorm.DB
}

func NewPostgresDB(db *gorm.DB) *pgDb {
	return &pgDb{
		db: db,
	}
}

func (pg pgDb) GetCustomer(id uint) (entities.Customer, error) {

	var dbCust db.Customer

	if res := pg.db.Take(&dbCust, id); res.Error != nil {
		return entities.Customer{}, res.Error
	}

	cust := entities.Customer{
		ID:           dbCust.ID,
		FirstName:    dbCust.FirstName,
		LastName:     dbCust.LastName,
		SignedWaiver: dbCust.SignedWaiver,
	}

	return cust, nil
}

func (pg pgDb) SaveCustomer(customerPtr *entities.Customer) error {
	customer := *customerPtr
	dbCust := db.Customer{
		FirstName:    customer.FirstName,
		LastName:     customer.LastName,
		SignedWaiver: customer.SignedWaiver,
	}
	if res := pg.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&dbCust); res.Error != nil {
		return res.Error
	}
	return nil
}
