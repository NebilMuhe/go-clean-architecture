package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

func SendSucessResponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Status: "ok",
		Data:   data,
	})
}
