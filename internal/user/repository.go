package user

import (
	"app/test/pkg/db"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	database *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (r *UserRepository) Create(user *User) (*User, error) {
	result := r.database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := r.database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*User, error) {
	var user User
	result := r.database.DB.First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) FindById(id uint) (*User, error) {
	var user User
	result := r.database.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Update(user *User) (*User, error) {
	result := r.database.DB.Clauses(clause.Returning{}).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepository) Delete(id uint) error {
	result := r.database.DB.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
