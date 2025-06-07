package service

import (
	"context"
	"todolist/internal/api"
	"todolist/internal/types"
	"todolist/pkg/logger"
)

type CategoryBiz interface {
	CreateCategory(ctx context.Context, domain *types.CategoryDomain) (int64, error)
	GetCategoryById(ctx context.Context, id int64) (*types.CategoryDomain, error)
}

func (service *CategoryService) CreateCategory(ctx context.Context, domain *types.CategoryDomain) (*types.CategoryDomain, error) {
	id, err := service.biz.CreateCategory(ctx, domain)
	if err != nil {
		return nil, err
	}

	domain.Id = id
	return domain, nil
}

type CategoryService struct {
	biz CategoryBiz
	log logger.Logger
}

func (service *CategoryService) GetCategoryById(ctx context.Context, id int64) (*types.CategoryDomain, error) {
	return service.biz.GetCategoryById(ctx, id)
}

func NewCategoryService(biz CategoryBiz, log logger.Logger) api.CategoryServiceInterface {
	return &CategoryService{
		biz: biz,
		log: log,
	}
}
