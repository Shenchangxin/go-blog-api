package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go-blog-api/user-web/global"
	"go-blog-api/user-web/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {
	consulConfig := global.ServerConfig.ConsulConfig
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.ServerConfig.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 [用户服务] 失败")
	}
	//TODO 后续用户服务信息变更需要维护该全局变量,添加grpc连接池功能
	userClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userClient

}

func InitSrvConn2() {
	//从注册中心获取用户服务信息
	cfg := api.DefaultConfig()
	consulConfig := global.ServerConfig.ConsulConfig
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	userSrvHost := ""
	userSrvPort := 0

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 [用户服务] 失败")
		return
	}
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorf("用户服务连接失败：%s", err.Error())
	}
	//TODO 后续用户服务信息变更需要维护该全局变量,添加grpc连接池功能
	userClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userClient
}
