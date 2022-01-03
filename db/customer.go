package db

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName    string
	LastName     string
	SignedWaiver bool

	Appointments []Appointment
}
