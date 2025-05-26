package dao

import (
	"context"
	biz "todolist/internal/biz"

	"gorm.io/gorm"
)

type GORMUserDao struct {
	db *gorm.DB
}

func (ud *GORMUserDao) FindByPasswordAndEmail(ctx context.Context, email, password string) (*biz.User, error) {
	var user biz.User
	err := ud.db.Where("email = ? and password = ?", email, password).First(&user).Error
	return &user, err
}

func (ud *GORMUserDao) Create(ctx context.Context, user *biz.User) (int64, error) {
	err := ud.db.Create(user).Error
	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

func NewGORMUserDao(db *gorm.DB) biz.UserDao {
	return &GORMUserDao{
		db: db,
	}
}

func (ud *GORMUserDao) FindById(ctx context.Context, id int64) (*biz.User, error) {
	var user biz.User
	err := ud.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
