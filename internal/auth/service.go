package auth

import (
	"app/test/internal/user"
	"errors"
)

type AuthService struct {
	userRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Register(email, username, password string) (string, error) {
	existedUser, _ := s.userRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	user := user.NewUser(email, username, password)
	user, err := s.userRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
