package usecase

import (
	"anon-chat/internal/domain"
	"anon-chat/internal/repo"
	"context"
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
