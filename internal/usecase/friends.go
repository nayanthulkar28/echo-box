package usecase

import (
	"context"
	"echo-box/internal/domain"
	"echo-box/internal/repo"
	"fmt"
)

type FriendUsecase struct {
	userRepo *repo.UserRepo
}

func NewFriendUsecase(ur *repo.UserRepo) *FriendUsecase {
	return &FriendUsecase{
		userRepo: ur,
	}
}

func (fu *FriendUsecase) MakeFriends(user1 domain.UserData, user2 domain.UserData) error {
	u1, err := fu.userRepo.GetUserByUsername(context.Background(), user1.Username)
	if err != nil {
		return err
	}
	u2, err := fu.userRepo.GetUserByUsername(context.Background(), user2.Username)
	if err != nil {
		return err
	}

	p1 := false
	p2 := false
	for _, f := range u1.Friends {
		if f == u2.Username {
			p1 = true
			break
		}
	}
	for _, f := range u2.Friends {
		if f == u1.Username {
			p2 = true
			break
		}
	}

	if !p1 {
		u1.Friends = append(u1.Friends, u2.Username)
	}
	if !p2 {
		u2.Friends = append(u2.Friends, u1.Username)
	}

	err = fu.userRepo.UpdateUser(context.Background(), u1)
	if err != nil {
		return err
	}
	err = fu.userRepo.UpdateUser(context.Background(), u2)
	return err
}

func (fu *FriendUsecase) GetFriendsByUsername(username string) ([]domain.UserData, error) {
	friendList := make([]domain.UserData, 0)
	friends, err := fu.userRepo.GetFriendsByUsername(context.Background(), username)
	if err != nil {
		return []domain.UserData{}, fmt.Errorf("unable to fetch friends %v", err)
	}

	for _, f := range friends {
		var fl domain.UserData
		fl.Username = f.Username
		friendList = append(friendList, fl)
	}
	return friendList, nil
}
