package router

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/user-web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	UserRouter.GET("list", api.GetUserList)
	UserRouter.POST("login", api.Login)
}
