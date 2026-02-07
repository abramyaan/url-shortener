package event

import "gorm.io/gorm"

type EventRepository struct {
	Db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{Db:db}
}

func (repo *EventRepository) Create (event Event) {
	repo.Db.Create(event)
}