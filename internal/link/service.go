package link

import (
	"url-shortener/internal/event" // Импортируем события
	"url-shortener/pkg/random"
)

type LinkService struct {
	Repo      *LinkRepository
	EventRepo *event.EventRepository // Добавляем репо событий
}

func NewLinkService(repo *LinkRepository, eventRepo *event.EventRepository) *LinkService {
	return &LinkService{
		Repo:      repo,
		EventRepo: eventRepo,
	}
}

func (s *LinkService) Create(url string, userID uint) (*Link, error) {
	hash := random.String(6)
	link := &Link{
		Url:    url,
		Hash:   hash,
		UserID: userID,
	}
	return s.Repo.Create(link)
}

func (s *LinkService) GetByHash(hash string) (*Link, error) {
	link, err := s.Repo.GetByHash(hash)
	if err != nil {
		return nil, err
	}

	// Когда ссылку нашли — создаем событие клика
	go s.EventRepo.Create(&event.Event{
		LinkID: link.ID,
	})

	return link, nil
}