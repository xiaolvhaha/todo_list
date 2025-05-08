package service

import (
	"context"
	"todolist/pkg/logger"
)

type UserDao interface {
	FindById(ctx context.Context, id int64) (UserDao, error)
}

type UserService struct {
	userDao UserDao
	l       logger.Logger
}

func NewUserService(userDao UserDao, l logger.Logger) *UserService {
	return &UserService{
		userDao: userDao,
		l:       l,
	}
}
