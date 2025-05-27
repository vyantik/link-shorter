package link

import "app/test/pkg/db"

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(db *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: db,
	}
}

func (r *LinkRepository) Create(link *Link) {

}
