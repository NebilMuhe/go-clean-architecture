package user

import (
	"go-clean-architecture/internal/constant/model/dto"
	"go-clean-architecture/internal/constant/model/response"
	"go-clean-architecture/internal/handler"
	"go-clean-architecture/internal/module"
	"go-clean-architecture/platform/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	userModule module.User
	log        logger.Logger
}

func Init(module module.User, log logger.Logger) handler.UserHandler {
	return &UserHandler{
		userModule: module,
		log:        log,
	}
}

// SignUp is a handler function to handle user sign up request
// @Summary Sign up a user
// @Description Sign up a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body dto.User true "User details"
// @Success 200 {object} dto.UserResponse
// @Router /user/signup [post]
func (u *UserHandler) SignUp(ctx *gin.Context) {
	var user dto.User
	if err := ctx.ShouldBind(&user); err != nil {
		u.log.Error(ctx, "failed to bind user details", zap.Error(err))
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	userResponse, err := u.userModule.SignUp(ctx, user)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}

	response.SendSucessResponse(ctx, http.StatusOK, userResponse)
}

// Login is a handler function to handle user login request
// @Summary Login a user
// @Description Login a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body dto.User true "User details"
// @Success 200 {object} dto.Token
// @Router /user/login [post]
// @Security Bearer Auth
func (u *UserHandler) Login(ctx *gin.Context) {
	var user dto.User

	if err := ctx.ShouldBind(&user); err != nil {
		u.log.Error(ctx, "failed to bind user details", zap.Error(err))
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}

	token, err := u.userModule.Login(ctx, user)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}

	response.SendSucessResponse(ctx, http.StatusOK, token)
}
