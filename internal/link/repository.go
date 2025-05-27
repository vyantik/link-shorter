package link

import (
	"app/test/pkg/db"
	"log"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(db *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: db,
	}
}

func (r *LinkRepository) Create(link *Link) (*Link, error) {
	result := r.Database.DB.Create(link)
	if result.Error != nil {
		log.Println("[Link] - [Repository] - [ERROR] : error creating link")
		return nil, result.Error
	}
	return link, nil
}

func (r *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := r.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		log.Println("[Link] - [Repository] - [ERROR] : error getting link by hash")
		return nil, result.Error
	}
	return &link, nil
}
