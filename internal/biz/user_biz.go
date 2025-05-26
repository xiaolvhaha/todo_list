package biz

import (
	"context"
	"todolist/internal/service"
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
	FindByPasswordAndEmail(ctx context.Context, email, password string) (*User, error)
	Create(ctx context.Context, user *User) (int64, error)
}

type UserUsecase struct {
	userDao UserDao
	l       logger.Logger
}

func (u *UserUsecase) FindByPasswordAndEmail(ctx context.Context, email, password string) (*types.UserDomain, error) {
	user, err := u.userDao.FindByPasswordAndEmail(ctx, email, password)
	if err != nil {
		return nil, err
	}
	return u.toDomain(user), nil
}

func (u *UserUsecase) CreateUser(ctx context.Context, user *types.UserDomain) (int64, error) {
	return u.userDao.Create(ctx, u.toModel(user))
}

func NewUserBiz(userDao UserDao, l logger.Logger) service.UserBiz {
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
