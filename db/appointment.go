package db

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	StartTime  *time.Time
	EndTime    *time.Time
	ServiceID  int
	Service    Service
	CustomerID uint
}
