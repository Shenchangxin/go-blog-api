package router

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/user-web/api"
	"go-blog-api/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
	UserRouter.POST("login", api.Login)
}
