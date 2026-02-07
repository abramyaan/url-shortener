package link

import "url-shortener/pkg/random"

type LinkService struct {
	Repo *LinkRepository
}

func NewLinkService(repo *LinkRepository) *LinkService {
	return &LinkService{Repo: repo}
}

func (s *LinkService) Create(url string, userID uint) (*Link, error) {
	hash := random.String(6)
	link := &Link {
		Url: url,
		Hash: hash,
		UserID: userID,
	}
	return s.Repo.Create(link)
}

func (s *LinkService) GetByHash(hash string) (*Link, error) {
	return s.Repo.GetByHash(hash)
}