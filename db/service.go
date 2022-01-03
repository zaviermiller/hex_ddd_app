package db

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Length int
}
