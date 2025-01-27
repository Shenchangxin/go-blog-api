package initialize

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/user-web/middlewares"
	router "go-blog-api/user-web/router"
	"go.uber.org/zap"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors()) //配置跨域
	ApiGroup := Router.Group("/v1")
	zap.S().Info("----配置用户相关URL----")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	return Router
}
