package dao

import (
	"context"
	biz "todolist/internal/biz"

	"gorm.io/gorm"
)

type GORMUserDao struct {
	db *gorm.DB
}

func NewGORMUserDao(db *gorm.DB) *GORMUserDao {
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
