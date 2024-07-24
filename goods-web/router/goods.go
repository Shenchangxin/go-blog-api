package router

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/goods-web/api/goods"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("list", goods.List)
		GoodsRouter.GET("createGoods", goods.CreateGoods)
	}
}
