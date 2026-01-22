// Package services
package services

import (
	"errors"

	"gorm.io/gorm"

	"backend/model"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

func (s *AuthService) Register(email, password string) error {

	user := model.User{

		Email:    email,
		Password: password,
	}

	result := s.DB.Create(&user)

	if result.Error != nil {
		return errors.New("username already taken")
	}

	return nil

}

func (s *AuthService) Login(email, password string) (*model.User, error) {

	var user model.User

	result := s.DB.Where("email = ? AND password = ?", email, password).First(&user)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	if user.Password != password {

		return nil, errors.New("invalid password")

	}

	return &user, nil
}
