package link

import (
	"url-shortener/internal/event"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url string `json:"url" gorm:"uniqueIndex;not null"`
	Hash string `json:"hash" gorm:"uniqueIndex;not null"`
	UserID uint
	Events []event.Event `gorm:"foreignKey:LinkID"`
}