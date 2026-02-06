package link

import "gorm.io/gorm"

type LinkRepository struct {
	Db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{
		Db: db,
	}
}

func (repo *LinkRepository) Create(link *Link)(*Link, error) {
	result:=repo.Db.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result:=repo.Db.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}