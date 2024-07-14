package initialize

import (
	"github.com/gin-gonic/gin"
	userRouter "go-blog-api/user-web/router"
	"go.uber.org/zap"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/v1")
	zap.S().Info("----配置用户相关URL----")
	userRouter.InitUserRouter(ApiGroup)
	return Router
}
