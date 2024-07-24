package config

type GoodsServerInfo struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      uint64 `mapstructure:"port" json:"port"`
	NameSpace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataId    string `mapstructure:"dataId" json:"dataId"`
	Group     string `mapstructure:"group" json:"group"`
}

type ServerConfig struct {
	Name            string          `mapstructure:"name" json:"name"`
	Host            string          `mapstructure:"host" json:"host"`
	Port            int             `mapstructure:"port" json:"port"`
	Tags            []string        `mapstructure:"tags" json:"tags"`
	GoodsServerInfo GoodsServerInfo `mapstructure:"goods_srv" json:"goods_srv"`
	ConsulConfig    ConsulConfig    `mapstructure:"consul" json:"consul"`
	RedisConfig     RedisConfig     `mapstructure:"redis" json:"redis"`
	NacosConfig     NacosConfig     `mapstructure:"nacos" json:"nacos"`
}
