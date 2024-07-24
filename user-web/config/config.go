package config

type UserServerInfo struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"host"`
}

type AliSmsConfig struct {
	ApiKey    string `mapstructure:"key" json:"host"`
	ApiSecret string `mapstructure:"secret" json:"secret"`
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
	Name         string         `mapstructure:"name" json:"name"`
	Host         string         `mapstructure:"host" json:"host"`
	Tags         []string       `mapstructure:"tags" json:"tags"`
	Port         int            `mapstructure:"port" json:"port"`
	UserSrvInfo  UserServerInfo `mapstructure:"user_srv" json:"user_srv"`
	JWTConfig    JWTConfig      `mapstructure:"jwt" json:"jwt"`
	AliSmsConfig AliSmsConfig   `mapstructure:"ali-sms" json:"ali-sms"`
	ConsulConfig ConsulConfig   `mapstructure:"consul" json:"consul"`
	RedisConfig  RedisConfig    `mapstructure:"redis" json:"redis"`
	NacosConfig  NacosConfig    `mapstructure:"nacos" json:"nacos"`
}
