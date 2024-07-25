package initialize

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/goods-web/middlewares"
	router "go-blog-api/goods-web/router"
	"go.uber.org/zap"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors()) //配置跨域
	ApiGroup := Router.Group("/v1")
	zap.S().Info("----配置商品相关URL----")
	router.InitGoodsRouter(ApiGroup)
	router.InitCategoryRouter(ApiGroup)
	router.InitBannerRouter(ApiGroup)
	router.InitBrandRouter(ApiGroup)
	return Router
}
