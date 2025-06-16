package dao

import (
	"context"
	"gorm.io/gorm"
	"todolist/internal/biz"
	"todolist/pkg/logger"
)

type GORMTaskDao struct {
	db  *gorm.DB
	log logger.Logger
}

func (dao *GORMTaskDao) GetList(ctx context.Context, limit, offset int64) ([]*biz.Task, error) {
	var list []*biz.Task
	err := dao.db.Limit(int(limit)).Offset(int(offset)).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, err
}

func (dao *GORMTaskDao) CreateTask(ctx context.Context, task *biz.Task) (int64, error) {
	err := dao.db.Create(task).Error
	if err != nil {
		return 0, err
	}

	return task.Id, nil
}

func NewGORMTaskDao(db *gorm.DB, log logger.Logger) biz.TaskDao {
	return &GORMTaskDao{
		db:  db,
		log: log,
	}
}
