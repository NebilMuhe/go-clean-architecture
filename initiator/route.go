package initiator

import (
	"go-clean-architecture/internal/glue/routing/user"

	"github.com/gin-gonic/gin"
)


func InitRoute(group *gin.RouterGroup,handler Handler){
	user.InitRoute(group,handler.user)
}