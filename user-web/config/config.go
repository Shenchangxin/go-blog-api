package config

type UserServerInfo struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}
type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type AliSmsConfig struct {
	ApiKey    string `mapstructure:"key"`
	ApiSecret string `mapstructure:"secret"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Name         string         `mapstructure:"name"`
	Port         int            `mapstructure:"port"`
	UserSrvInfo  UserServerInfo `mapstructure:"user_srv"`
	JWTConfig    JWTConfig      `mapstructure:"jwt"`
	AliSmsConfig AliSmsConfig   `mapstructure:"ali-sms"`
	ConsulConfig ConsulConfig   `mapstructure:"consul"`
	RedisConfig  RedisConfig    `mapstructure:"redis"`
}
