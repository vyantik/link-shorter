package link

import (
	"app/test/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	database *db.Db
}

func NewLinkRepository(db *db.Db) *LinkRepository {
	return &LinkRepository{
		database: db,
	}
}

func (r *LinkRepository) Create(link *Link) (*Link, error) {
	result := r.database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (r *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := r.database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (r *LinkRepository) Update(link *Link) (*Link, error) {
	result := r.database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (r *LinkRepository) Delete(id uint) error {
	result := r.database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *LinkRepository) GetById(id uint) (*Link, error) {
	var link Link
	result := r.database.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (r *LinkRepository) GetCount(email string) int64 {
	var count int64
	r.database.
		Table("links").
		Where("deleted_at IS NULL").
		Count(&count)
	return count
}

func (r *LinkRepository) GetAll(limit, offset int) []Link {
	var links []Link

	r.database.
		Table("links").
		Where("deleted_at IS NULL").
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Scan(&links)

	return links
}
