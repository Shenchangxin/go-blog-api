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
		GoodsRouter.GET("deleteGoods", goods.Delete)
		GoodsRouter.GET("updateGoods", goods.Update)
		GoodsRouter.GET("updateStatus", goods.UpdateStatus)
		GoodsRouter.GET("stocks", goods.Stocks)
	}
}
