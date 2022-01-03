package ports

import "hex_ddd_app/internal/client/entities"

// these interfaces produce side effects !! not sure how bad that is, seems fine here

// driven port - provided by external actor
type CustomerRepository interface {
	GetCustomer(id uint) (entities.Customer, error)
	SaveCustomer(customer *entities.Customer) error
}

// driver port - provided by core
type CustomerService interface {
	Get(id uint) (entities.Customer, error)
	Create(customer *entities.Customer) error
	// ScheduleAppointment() (entities.Customer, error)
}
