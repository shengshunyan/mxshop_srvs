package config

import "mxshop_srvs/common/config"

type ServerConfig struct {
	Name       string              `mapstructure:"name"`
	Host       string              `mapstructure:"host"`
	Port       int                 `mapstructure:"port"`
	MysqlInfo  config.MysqlConfig  `mapstructure:"mysql"`
	ConsulInfo config.ConsulConfig `mapstructure:"consul"`
	Redis      RedisConfig         `mapstructure:"redis"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
