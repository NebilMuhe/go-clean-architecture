package handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}
