package initialize

import (
	"mxshop_srvs/common/initialize"
	"mxshop_srvs/goods-srv/global"
)

// 服务注册
func InitRegister() {
	serverConfig := global.ServerConfig
	initialize.InitRegister(serverConfig)
}

func CloseRegister() {
	initialize.CloseRegister()
}
