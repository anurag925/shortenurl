package service

import (
	"context"
	"errors"
	"time"

	"github.com/anurag/shortenurl/internal/db/models"
	"github.com/anurag/shortenurl/internal/db/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, username, password string) (string, error)
	Login(ctx context.Context, username, password string) (string, error)
}

type authService struct {
	userRepo repositories.UserRepository
	secret   string
}

func NewAuthService(userRepo repositories.UserRepository, secret string) AuthService {
	return &authService{
		userRepo: userRepo,
		secret:   secret,
	}
}

func (s *authService) Register(ctx context.Context, username, password string) (string, error) {
	// Check if user exists
	_, err := s.userRepo.FindByUsername(ctx, username)
	if err == nil {
		return "", errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Create user
	user := &models.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return "", err
	}

	// Generate JWT token
	return s.generateToken(user.ID)
}

func (s *authService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return s.generateToken(user.ID)
}

func (s *authService) generateToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(s.secret))
}
