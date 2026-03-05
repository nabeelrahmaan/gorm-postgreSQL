package service

import (
	"authentication/models"
	"authentication/repository"
	"authentication/utils"
	"errors"
	"time"
)

type AuthService struct {
	repo *repository.UserRepository
}

func (s *AuthService) Signup(name, email, password string) error {

	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: hash,
	}

	return s.repo.Create(&user)
}

func (s *AuthService) Login(email, password string) (string, string, error) {

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", "", errors.New("invalid credentials")
	}

	access, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", errors.New("failed to create access token")
	}

	refresh, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", errors.New("failed to create refresh token")
	}

	ref := models.RefreshToken{
		UserID:    user.ID,
		Token:     refresh,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = s.repo.CreateRefresh(&ref)
	if err != nil {
		return "", "", errors.New("failed to store refresh token")
	}

	return access, refresh, nil
}
