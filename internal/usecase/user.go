package usecase

import (
	"context"
	"echo-box/internal/domain"
	"echo-box/internal/repo"
)

type UserUsecase struct {
	userRepo *repo.UserRepo
}

func NewUserUsecase(ur *repo.UserRepo) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

func (uu *UserUsecase) GetUser(ctx context.Context) (domain.UserResponse, error) {
	user := ctx.Value("user").(*domain.User)
	return domain.UserResponse{User: domain.UserData{Username: user.Username}}, nil
}
