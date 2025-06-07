package dao

import (
	"context"
	"gorm.io/gorm"
	"todolist/internal/biz"
	"todolist/pkg/logger"
)

type GORMCategoryDao struct {
	db  *gorm.DB
	log logger.Logger
}

func (dao *GORMCategoryDao) GetCategoryById(ctx context.Context, id int64) (*biz.Category, error) {
	var category biz.Category
	err := dao.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (dao *GORMCategoryDao) CreateCategory(ctx context.Context, category *biz.Category) (int64, error) {
	err := dao.db.Create(category).Error
	if err != nil {
		return 0, err
	}

	return category.Id, nil
}

func NewGORMCategoryDao(db *gorm.DB, log logger.Logger) biz.CategoryDao {
	return &GORMCategoryDao{
		db:  db,
		log: log,
	}
}
