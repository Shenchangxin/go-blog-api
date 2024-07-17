package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-blog-api/user-web/global"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {

	debug := GetEnvInfo("GO-BLOG")
	configFilePrefix := "config"

	configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
	if !debug {
		configFileName = fmt.Sprintf("user-web/%s-dev.yaml", configFilePrefix)
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息：%v", global.ServerConfig)
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		zap.S().Infof("检测到配置信息发生变化，配置信息：%v", global.ServerConfig)
	})

}
