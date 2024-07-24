package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go-blog-api/goods-web/global"
	"go-blog-api/goods-web/initialize"
	"go-blog-api/goods-web/utils"
	"go-blog-api/goods-web/utils/register/consul"
	myvalidator "go-blog-api/goods-web/validators"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
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
	registryClient := consul.NewRegistryClient(global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port)
	uuidNew := uuid.NewV4()
	serviceId := fmt.Sprintf("%s", uuidNew)
	err := registryClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panicf("注册商品服务：%d 失败", global.ServerConfig.Port)
		panic(err)
	}

	zap.S().Infof("启动服务器：%d", global.ServerConfig.Port)

	if err := Routers.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panicf("启动服务器：%d 失败", global.ServerConfig.Port)
	}
	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = registryClient.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败:", err.Error())
	} else {
		zap.S().Info("注销成功:")
	}
}
