package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex"`
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"password"`
}

func NewUser(email, username, password string) *User {
	return &User{
		Email:    email,
		Username: username,
		Password: password,
	}
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
