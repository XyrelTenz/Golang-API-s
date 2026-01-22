// Package services
package services

import (
	"errors"
	"sync"

	"backend/model"
)

type AuthService struct {
	mu    sync.RWMutex
	users map[string]model.User
}

func NewAuthService() *AuthService {
	return &AuthService{
		users: make(map[string]model.User),
	}
}

func (s *AuthService) Register(username, password string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.users[username]; exist {
		return errors.New("users already exist")

	}

	newUser := &model.User{
		ID:       len(s.users) + 1,
		Username: username,
		Password: password,
	}

	s.users[username] = *newUser

	return nil

}

func (s *AuthService) Login(username, password string) bool {

	s.mu.Lock()
	defer s.mu.Unlock()

	storedPassword, exist := s.users[password]

	if !exist {
		return false
	}

	return storedPassword == s.users[password]
}

func (s *AuthService) GetProfile(username string) (model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exist := s.users[username]

	if !exist {

		return model.User{}, errors.New("user not found")

	}

	return user, nil

}
