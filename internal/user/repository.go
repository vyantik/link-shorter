package user

import (
	"app/test/pkg/db"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{
		Database: db,
	}
}

func (r *UserRepository) Create(user *User) (*User, error) {
	result := r.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*User, error) {
	var user User
	result := r.Database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetByUsername(username string) (*User, error) {
	var user User
	result := r.Database.DB.First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetById(id uint) (*User, error) {
	var user User
	result := r.Database.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Update(user *User) (*User, error) {
	result := r.Database.DB.Clauses(clause.Returning{}).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepository) Delete(id uint) error {
	result := r.Database.DB.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
