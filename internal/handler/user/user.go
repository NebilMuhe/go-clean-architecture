package user

import (
	"go-clean-architecture/internal/constant/model/dto"
	"go-clean-architecture/internal/handler"
	"go-clean-architecture/internal/module"
	"go-clean-architecture/platform/logger"

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


func (u *UserHandler) SignUp(ctx *gin.Context) {
	var user dto.User
	if err := ctx.ShouldBind(&user); err != nil{
		u.log.Error(ctx,"failed to bind user details",zap.Error(err))
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}
	u.userModule.SignUp(ctx,user)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var user dto.User

	if err := ctx.ShouldBind(&user); err != nil {
		u.log.Error(ctx,"failed to bind user details",zap.Error(err))
		_ = ctx.Error(err)
		ctx.Abort()
		return
	}

	u.userModule.Login(ctx,user)
}



