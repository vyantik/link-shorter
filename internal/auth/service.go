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

	newUser, err := user.NewUser(email, username, password)
	if err != nil {
		return "", err
	}

	createdUser, err := s.userRepository.Create(newUser)
	if err != nil {
		return "", err
	}
	return createdUser.Email, nil
}
