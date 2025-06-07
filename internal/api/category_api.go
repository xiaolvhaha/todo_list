package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/internal/types"
	"todolist/internal/types/request"
	"todolist/pkg/logger"
)

type CategoryServiceInterface interface {
	CreateCategory(ctx context.Context, domain *types.CategoryDomain) (*types.CategoryDomain, error)
	GetCategoryById(ctx context.Context, id int64) (*types.CategoryDomain, error)
}

type CategoryApi struct {
	service CategoryServiceInterface
	log     logger.Logger
}

func (api *CategoryApi) RegisterRouter(engin *gin.Engine) {
	group := engin.Group("/category")
	group.POST("/create", api.Create)
}

func (api *CategoryApi) Create(ctx *gin.Context) {
	var req request.CreateCategoryRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.Result{
			Code: 4,
			Msg:  "invalid request body",
			Data: nil,
		})
	}

	category, err := api.service.CreateCategory(ctx, &types.CategoryDomain{
		Name: req.Name,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.Result{
			Code: 5,
			Msg:  "internal server error",
			Data: nil,
		})
	}

	ctx.JSON(http.StatusOK, types.Result{
		Code: 0,
		Data: category,
		Msg:  "success",
	})

	return
}

func NewCategoryApi(service CategoryServiceInterface, log logger.Logger) *CategoryApi {
	return &CategoryApi{
		service: service,
		log:     log,
	}
}
