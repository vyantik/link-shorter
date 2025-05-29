package auth_test

import (
	"app/test/internal/auth"
	"app/test/internal/user"
	"testing"
)

const (
	initialEmail    = "test@test.com"
	initialUsername = "test"
	initialPassword = "12345678"
)

type MockUserRepository struct {
}

func (m *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: initialEmail,
	}, nil
}

func (m *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, initialUsername, initialPassword)
	if err != nil {
		t.Errorf("Error registering user: %v", err)
	}
	if email != initialEmail {
		t.Errorf("Email is not correct: %v", email)
	}
}
