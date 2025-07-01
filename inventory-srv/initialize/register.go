package initialize

import (
	"mxshop_srvs/common/initialize"
	"mxshop_srvs/inventory-srv/global"
)

// 服务注册
func InitRegister() {
	serverConfig := global.ServerConfig
	initialize.InitRegister(serverConfig)
}

func CloseRegister() {
	initialize.CloseRegister()
}
