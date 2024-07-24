package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go-blog-api/goods-web/global"
	"go-blog-api/goods-web/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {
	consulConfig := global.ServerConfig.ConsulConfig
	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.ServerConfig.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 [用户服务] 失败")
	}
	//TODO 后续用户服务信息变更需要维护该全局变量,添加grpc连接池功能
	goodsClient := proto.NewGoodsClient(goodsConn)
	global.GoodsSrvClient = goodsClient

}

func InitSrvConn2() {
	//从注册中心获取用户服务信息
	cfg := api.DefaultConfig()
	consulConfig := global.ServerConfig.ConsulConfig
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	goodsSrvHost := ""
	goodsSrvPort := 0

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.GoodsServerInfo.Name))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		goodsSrvHost = value.Address
		goodsSrvPort = value.Port
		break
	}
	if goodsSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 [用户服务] 失败")
		return
	}
	goodsConn, err := grpc.Dial(fmt.Sprintf("%s:%d", goodsSrvHost, goodsSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorf("用户服务连接失败：%s", err.Error())
	}
	//TODO 后续用户服务信息变更需要维护该全局变量,添加grpc连接池功能
	goodsClient := proto.NewGoodsClient(goodsConn)
	global.GoodsSrvClient = goodsClient
}
