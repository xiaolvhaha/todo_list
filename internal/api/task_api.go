package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"todolist/internal/types"
	"todolist/internal/types/request"
	"todolist/internal/types/response"
	"todolist/pkg/logger"
)

type TaskServiceInterface interface {
	CreateTask(ctx context.Context, task *types.TaskDomain) (int64, error)
	GetTaskList(ctx context.Context, page int64) ([]*types.TaskDomain, error)
}

type TaskApi struct {
	service         TaskServiceInterface
	categoryService CategoryServiceInterface
	log             logger.Logger
}

func NewTaskApi(service TaskServiceInterface, categoryService CategoryServiceInterface, log logger.Logger) *TaskApi {
	return &TaskApi{
		service:         service,
		categoryService: categoryService,
		log:             log,
	}
}

func (api *TaskApi) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/task")
	group.POST("/create", api.CreateTask)
	group.GET("/list", api.GetList)
}

func (api *TaskApi) CreateTask(ctx *gin.Context) {
	var req request.CreateTaskRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.Result{
			Code: 4,
			Msg:  "invalid param",
			Data: nil,
		})

		return
	}

	category, err := api.categoryService.GetCategoryById(ctx, req.CategoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.Result{
			Code: 4,
			Msg:  "invalid param",
			Data: nil,
		})

		return
	}

	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, req.Deadline)
	if err != nil {
		api.log.Error("parse time error", logger.Field{
			Key:   "error",
			Value: err.Error(),
		})
		ctx.JSON(http.StatusBadRequest, types.Result{
			Code: 4,
			Msg:  "invalid param",
			Data: nil,
		})

		return
	}

	// 转换为 int64 毫秒级时间戳
	timestampMs := t.UnixMilli()

	id, err := api.service.CreateTask(ctx, &types.TaskDomain{
		Title:    req.Title,
		Desc:     req.Desc,
		Property: req.Priority,
		Category: types.CategoryDomain{Id: category.Id},
		Status:   0,
		Deadline: timestampMs,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.Result{
			Code: 5,
			Msg:  "internal server error",
			Data: nil,
		})

		return
	}

	domain := types.TaskDomain{
		Id:       id,
		Title:    req.Title,
		Desc:     req.Desc,
		Property: req.Priority,
		Deadline: timestampMs,
		Status:   0,
		Category: types.CategoryDomain{
			Id:   category.Id,
			Name: category.Name,
		},
	}

	ctx.JSON(http.StatusOK, types.Result{
		Code: 0,
		Msg:  "success",
		Data: response.TaskInfoResponse{
			Deadline: req.Deadline,
			Info:     domain,
		},
	})

}

func (api *TaskApi) GetList(ctx *gin.Context) {
	var req request.GetTaskListRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.Result{
			Code: 4,
			Msg:  "invalid param",
		})
	}

	list, err := api.service.GetTaskList(ctx, req.Page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.Result{
			Code: 5,
			Msg:  "internal server error",
		})
	}

	ctx.JSON(http.StatusOK, types.Result{
		Code: 0,
		Msg:  "success",
		Data: list,
	})

	return
}
