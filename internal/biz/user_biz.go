package biz

import (
	"context"
	service "todolist/internal/service"
	"todolist/internal/types"
	"todolist/pkg/logger"
)

type User struct {
	Id       int64  `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}

type UserDao interface {
	FindById(ctx context.Context, id int64) (*User, error)
}

type UserUsecase struct {
	userDao UserDao
	l       logger.Logger
}

func NewUserBiz(userDao UserDao, l logger.Logger) *UserUsecase {
	return &UserUsecase{
		userDao: userDao,
		l:       l,
	}
}

// FindById implements service.UserBiz.
func (u *UserUsecase) FindById(ctx context.Context, id int64) (*types.UserDomain, error) {
	user, err := u.userDao.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return u.toDomain(user), nil
}

func (u *UserUsecase) toModel(domain *types.UserDomain) *User {
	return &User{
		Id:       domain.Id,
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
	}
}

func (u *UserUsecase) toDomain(model *User) *types.UserDomain {
	return &types.UserDomain{
		Id:       model.Id,
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	}
}
