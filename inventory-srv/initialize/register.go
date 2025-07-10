package initialize

import (
	"mxshop_srvs/common/config"
	"mxshop_srvs/common/initialize"
	"mxshop_srvs/inventory-srv/global"
)

// 服务注册
func InitRegister() {
	serverConfig := global.ServerConfig
	initialize.InitRegister(&config.ServerConfig{
		Name:       serverConfig.Name,
		Host:       serverConfig.Host,
		Port:       serverConfig.Port,
		ConsulInfo: serverConfig.ConsulInfo,
	})
}

func CloseRegister() {
	initialize.CloseRegister()
}
