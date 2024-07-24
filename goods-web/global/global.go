package global

import (
	ut "github.com/go-playground/universal-translator"
	"go-blog-api/goods-web/config"
	"go-blog-api/goods-web/proto"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig    *config.NacosConfig = &config.NacosConfig{}
	Trans          ut.Translator
	GoodsSrvClient proto.GoodsClient
)
