package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"time"
	"todolist/internal/types"
	"todolist/internal/types/request"
	"todolist/internal/types/response"
	"todolist/pkg/logger"
)

func (u *UserApi) RegisterUserRouter(engine *gin.Engine) *gin.Engine {
	group := engine.Group("user")
	group.GET("/:id", u.GetUserById)
	group.POST("/register", u.Register)
	group.POST("/login", u.Login)
	return engine
}

func (u *UserApi) GetUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	byId, err := u.us.GetUserById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": byId})
}

func (u *UserApi) Register(ctx *gin.Context) {
	var req request.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.Result{
			Code: 4,
			Msg:  "INVALID PARAM",
			Data: nil,
		})
	}
	id, err := u.us.CreateUser(ctx, &types.UserDomain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.Result{
			Code: 5,
			Msg:  "INTERNAL ERROR",
			Data: nil,
		})
	}

	ctx.JSON(http.StatusOK, types.Result{
		Code: 0,
		Msg:  "SUCCESS",
		Data: response.UserInfo{
			Id:    id,
			Name:  req.Name,
			Email: req.Email,
		},
	})
	return
}

func (u *UserApi) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.Result{
			Code: 4,
			Msg:  "INVALID PARAM",
			Data: nil,
		})
		return
	}

	user, err := u.us.GetUserByPassAndEmail(ctx, req.Password, req.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.Result{})
		return
	}

	claims := types.UserClaim{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(types.UserSignKey))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.Result{})
		return
	}

	ctx.JSON(http.StatusOK, types.Result{
		Code: 0,
		Msg:  "SUCCESS",
		Data: response.LoginResponse{
			User: response.UserInfo{
				Id:    user.Id,
				Name:  user.Name,
				Email: user.Email,
			},
			Token: ss,
		},
	})

	return
}

type UserServiceInterface interface {
	GetUserById(context context.Context, id int64) (*types.UserDomain, error)
	GetUserByPassAndEmail(context context.Context, password, email string) (*types.UserDomain, error)
	CreateUser(context context.Context, user *types.UserDomain) (int64, error)
}

type UserApi struct {
	us  UserServiceInterface
	log logger.Logger
}

func NewUserApi(us UserServiceInterface, log logger.Logger) *UserApi {
	return &UserApi{
		us:  us,
		log: log,
	}
}
