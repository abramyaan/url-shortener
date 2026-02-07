package event

import "gorm.io/gorm"



type Event struct {
	gorm.Model
	LinkID uint
	UserID uint
}

