package usecase

import (
	"context"
	"echo-box/internal/domain"
	"echo-box/internal/repo"
	"echo-box/pkg"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	userRepo *repo.UserRepo
}

func NewAuthUsecase(ar *repo.UserRepo) *AuthUsecase {
	return &AuthUsecase{
		userRepo: ar,
	}
}

func (au *AuthUsecase) SignUp(ctx context.Context, req domain.SignUpRequest) (*domain.SignUpResponse, error) {
	user, err := au.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	if user.Id != "" {
		return nil, fmt.Errorf("username already exist")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, err
	}

	user.Id = uuid.NewString()
	user.Username = req.Username
	user.Password = string(hashPassword)
	user.Friends = make([]string, 0)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = au.userRepo.CreateUser(ctx, *user)
	if err != nil {
		return nil, fmt.Errorf("fail to create user")
	}
	return &domain.SignUpResponse{Status: "signed up"}, nil
}

func (au *AuthUsecase) Login(ctx context.Context, req domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := au.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid password %w", err)
	}

	token, err := pkg.CreateToken(req.Username)
	if err != nil {
		return nil, err
	}
	return &domain.LoginResponse{Status: "logged in",
		User:  domain.UserData{Username: user.Username},
		Token: token}, nil
}
