package service

import (
	"authentication/models"
	"authentication/repository"
	"authentication/utils"
	"errors"
	"fmt"
	"time"
)

type AuthService struct {
	Repo *repository.UserRepository
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
		Role: "user",
	}

	return s.Repo.Create(&user)
}

func (s *AuthService) Login(email, password string) (string, string, error) {

	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", "", errors.New("invalid credentials")
	}

	access, err := utils.GenerateAccessToken(user)
	fmt.Println("ACCESS ERROR:", err)
	if err != nil {
		return "", "", errors.New("failed to create access token")
	}

	refresh, err := utils.GenerateRefreshToken(user)
	if err != nil {
		return "", "", errors.New("failed to create refresh token")
	}

	ref := models.RefreshToken{
		UserID:    user.ID,
		Token:     refresh,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = s.Repo.CreateRefresh(&ref)
	if err != nil {
		return "", "", errors.New("failed to store refresh token")
	}

	return access, refresh, nil
}
