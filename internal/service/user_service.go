package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"todolist/internal/api"
	types "todolist/internal/types"
	"todolist/pkg/logger"
)

type UserBiz interface {
	FindById(ctx context.Context, id int64) (*types.UserDomain, error)
	CreateUser(ctx context.Context, user *types.UserDomain) (int64, error)
	FindByPasswordAndEmail(ctx context.Context, email, password string) (*types.UserDomain, error)
}

type UserService struct {
	uBiz UserBiz
	log  logger.Logger
}

func (us *UserService) GetUserByPassAndEmail(ctx context.Context, password, email string) (*types.UserDomain, error) {
	hash := md5.Sum([]byte(password))

	pass := hex.EncodeToString(hash[:])
	return us.uBiz.FindByPasswordAndEmail(ctx, email, pass)
}

func (us *UserService) CreateUser(context context.Context, user *types.UserDomain) (int64, error) {
	hash := md5.Sum([]byte(user.Password))

	password := hex.EncodeToString(hash[:])
	user.Password = password

	return us.uBiz.CreateUser(context, user)
}

func NewUserService(uBiz UserBiz, log logger.Logger) api.UserServiceInterface {
	return &UserService{
		uBiz: uBiz,
		log:  log,
	}
}

func (us *UserService) GetUserById(ctx context.Context, id int64) (*types.UserDomain, error) {
	return us.uBiz.FindById(ctx, id)
}
