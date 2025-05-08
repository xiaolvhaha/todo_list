package dao

import (
	"context"
	"gorm.io/gorm"
	"todolist/internal/service"
)

type User struct {
	id       int64  `gorm:"primaryKey"`
	name     string `gorm:"type:varchar(255)"`
	email    string `gorm:"type:varchar(255)"`
	password string `gorm:"type:varchar(255)"`
}

type GORMUserDao struct {
	db *gorm.DB
}

func NewGORMUserDao(db *gorm.DB) service.UserDao {
	return &GORMUserDao{
		db: db,
	}
}

func (G GORMUserDao) FindById(ctx context.Context, id int64) (service.UserDao, error) {
	//TODO implement me
	panic("implement me")
}

