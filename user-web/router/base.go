package router

import (
	"github.com/gin-gonic/gin"
	"go-blog-api/user-web/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		BaseRouter.POST("send_sms", api.SendSMS)
	}
}
