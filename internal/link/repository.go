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

func (repo *LinkRepository) GetByUserID(userID uint) ([]Link, error) {
	var links []Link
	err := repo.Db.Preload("Events").Where("user_id = ?", userID).Find(&links).Error
	return links, err
}



func (repo *LinkRepository) Delete(id uint, userID uint) error {
	// Удаляем только если ID ссылки и ID пользователя совпадают (защита)
	result := repo.Db.Where("id = ? AND user_id = ?", id, userID).Delete(&Link{})
	return result.Error
}