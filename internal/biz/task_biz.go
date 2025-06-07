package biz

import (
	"context"
	"todolist/internal/service"
	"todolist/internal/types"
	"todolist/pkg/logger"
)

type Task struct {
	Id         int64  `gorm:"primaryKey"`
	Title      string `gorm:"type:varchar(255)"`
	Desc       string `gorm:"type:varchar(255)"`
	CategoryId int64  `gorm:"column:category_id;type:varchar(255)"`
	Property   int64  `gorm:"type:tinyint(3);default:0"`
	Deadline   int64  `gorm:"type:bigint(20);default:0"`
	Status     int64  `gorm:"type:tinyint(3);default:0"`
}

type TaskDao interface {
	CreateTask(ctx context.Context, task *Task) (int64, error)
}

type TaskUseCase struct {
	dao TaskDao
	log logger.Logger
}

func (t *TaskUseCase) CreateTask(ctx context.Context, task *types.TaskDomain) (int64, error) {
	return t.dao.CreateTask(ctx, t.toModel(task))
}

func NewTaskUseCase(dao TaskDao, log logger.Logger) service.TaskBiz {
	return &TaskUseCase{dao: dao, log: log}
}

func (biz *TaskUseCase) toModel(task *types.TaskDomain) *Task {
	return &Task{
		Id:         task.Id,
		Title:      task.Title,
		Desc:       task.Desc,
		CategoryId: task.Category.Id,
		Property:   task.Property,
		Deadline:   task.Deadline,
		Status:     task.Status,
	}
}

func (biz *TaskUseCase) toDomain(task *Task) *types.TaskDomain {
	return &types.TaskDomain{
		Id:       task.Id,
		Title:    task.Title,
		Desc:     task.Desc,
		Status:   task.Status,
		Deadline: task.Deadline,
		Property: task.Property,
		Category: types.CategoryDomain{Id: task.CategoryId},
	}
}
