package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
	"net/http"
	"time"
	"todolist/internal/types"
	"todolist/internal/types/request"
	"todolist/internal/types/response"
)

type TaskServiceInterface interface {
	CreateTask(ctx context.Context, task *types.TaskDomain) (int64, error)
}

type TaskApi struct {
	service         TaskServiceInterface
	categoryService CategoryServiceInterface
	log             logr.Logger
}

func NewTaskApi(service TaskServiceInterface, categoryService CategoryServiceInterface, log logr.Logger) *TaskApi {
	return &TaskApi{
		service:         service,
		categoryService: categoryService,
		log:             log,
	}
}

func (api *TaskApi) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/task")
	group.POST("/create", api.CreateTask)
}

func (api *TaskApi) CreateTask(ctx *gin.Context) {
	var req request.CreateTaskRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.Result{
			Code: 4,
			Msg:  "invalid param",
			Data: nil,
		})
	}

	category, err := api.categoryService.GetCategoryById(ctx, req.CategoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.Result{
			Code: 4,
			Msg:  "invalid param",
			Data: nil,
		})
	}

	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, req.Deadline)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.Result{
			Code: 4,
			Msg:  "invalid param",
			Data: nil,
		})
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

	return
}
