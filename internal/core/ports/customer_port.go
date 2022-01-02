package ports

import (
	"hex_ddd_app/internal/core/domain"
)

type CustomerRepository interface {
	GetCustomer(id string) (domain.Customer, error)
	SaveCustomer(customer domain.Customer) error
}

type CustomerService interface {
	Get(id string) (domain.Customer, error)
	Create(customer domain.Customer) (domain.Customer, error)
	// ScheduleAppointment() (domain.Customer, error)
}
