package biz

import (
	"context"
	"todolist/internal/service"
	"todolist/internal/types"
	"todolist/pkg/logger"
)

type Category struct {
	Id   int64  `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

type CategoryDao interface {
	CreateCategory(ctx context.Context, category *Category) (int64, error)
	GetCategoryById(ctx context.Context, id int64) (*Category, error)
}

type CategoryUsecase struct {
	dao CategoryDao
	log logger.Logger
}

func (biz *CategoryUsecase) GetCategoryById(ctx context.Context, id int64) (*types.CategoryDomain, error) {
	byId, err := biz.dao.GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}

	return biz.toDomain(byId), nil
}

func NewCategoryBiz(dao CategoryDao, logger logger.Logger) service.CategoryBiz {
	return &CategoryUsecase{
		dao: dao,
		log: logger,
	}
}

func (biz *CategoryUsecase) CreateCategory(ctx context.Context, domain *types.CategoryDomain) (int64, error) {
	return biz.dao.CreateCategory(ctx, biz.toModel(domain))
}

func (biz *CategoryUsecase) toModel(domain *types.CategoryDomain) *Category {
	return &Category{
		Id:   domain.Id,
		Name: domain.Name,
	}
}

func (biz *CategoryUsecase) toDomain(model *Category) *types.CategoryDomain {
	return &types.CategoryDomain{
		Id:   model.Id,
		Name: model.Name,
	}
}
