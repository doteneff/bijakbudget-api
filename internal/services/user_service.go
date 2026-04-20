package services

import (
	"context"
	"errors"

	"github.com/doteneff/bijakbudget-api/internal/models"
	"github.com/doteneff/bijakbudget-api/internal/repositories"
	"github.com/doteneff/bijakbudget-api/internal/utils"
	"google.golang.org/api/idtoken"
)

type UserService interface {
	RegisterUser(user *models.User) (*models.AuthResponse, error)
	LoginUser(email, password string) (*models.AuthResponse, error)
	LoginSSOGoogle(token string) (*models.AuthResponse, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) RegisterUser(user *models.User) (*models.AuthResponse, error) {
	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	user.Provider = "local"

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	tokenStr, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{Token: tokenStr, User: *user}, nil
}

func (s *userService) LoginUser(email, password string) (*models.AuthResponse, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if user.Provider != "local" {
		return nil, errors.New("please login using your provider: " + user.Provider)
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	tokenStr, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{Token: tokenStr, User: *user}, nil
}

func (s *userService) LoginSSOGoogle(token string) (*models.AuthResponse, error) {
	// You may need to provide client ID if checking audience
	payload, err := idtoken.Validate(context.Background(), token, "")
	if err != nil {
		return nil, errors.New("invalid google token: " + err.Error())
	}

	email, ok := payload.Claims["email"].(string)
	if !ok {
		return nil, errors.New("no email in token")
	}

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		// User does not exist, register them
		name, _ := payload.Claims["name"].(string)
		
		user = &models.User{
			Email:    email,
			Name:     name,
			Provider: "google",
			Role:     "member",
		}
		
		if err := s.repo.Create(user); err != nil {
			return nil, err
		}
	}

	tokenStr, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{Token: tokenStr, User: *user}, nil
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}
