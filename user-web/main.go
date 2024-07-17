package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go-blog-api/user-web/global"
	"go-blog-api/user-web/initialize"
	"go-blog-api/user-web/utils"
	myvalidator "go-blog-api/user-web/validators"
	"go.uber.org/zap"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	initialize.InitSrvConn()

	viper.AutomaticEnv()
	dev := viper.GetBool("GO-BLOG")
	if dev {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	//注册验证器和翻译器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})

	}

	Routers := initialize.Routers()

	zap.S().Infof("启动服务器：%d", global.ServerConfig.Port)

	if err := Routers.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panicf("启动服务器：%d 失败", global.ServerConfig.Port)
	}

}
