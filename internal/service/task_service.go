package service

import (
	"context"
	"todolist/internal/api"
	"todolist/internal/types"
	"todolist/pkg/logger"
)

type TaskBiz interface {
	CreateTask(ctx context.Context, task *types.TaskDomain) (int64, error)
	GetTaskList(ctx context.Context, page int64) ([]*types.TaskDomain, error)
}

type TaskService struct {
	biz TaskBiz
	log logger.Logger
}

func (service *TaskService) GetTaskList(ctx context.Context, page int64) ([]*types.TaskDomain, error) {
	return service.biz.GetTaskList(ctx, page)
}

func (service *TaskService) CreateTask(ctx context.Context, task *types.TaskDomain) (int64, error) {
	return service.biz.CreateTask(ctx, task)
}

func NewTaskService(biz TaskBiz, log logger.Logger) api.TaskServiceInterface {
	return &TaskService{biz: biz, log: log}
}
