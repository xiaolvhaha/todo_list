package service

import (
	"context"
	v1 "todolist/api/user/v1"
	types "todolist/internal/types"
)

type UserBiz interface {
	FindById(ctx context.Context, id int64) (*types.UserDomain, error)
}

type UserService struct {
	uBiz UserBiz
}

func NewUserService(uBiz UserBiz) *UserService {
	return &UserService{
		uBiz: uBiz,
	}
}

func (us *UserService) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.UserInfoReply, error) {
	user, err := us.uBiz.FindById(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &v1.UserInfoReply{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
