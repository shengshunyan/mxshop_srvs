package initialize

import (
	"mxshop_srvs/common/initialize"
	"mxshop_srvs/common/utils"
	"mxshop_srvs/goods-srv/global"
)

func InitConfig() {
	env := utils.GetEnv()
	configFilePath := ""
	if env == "dev" {
		configFilePath = "goods-srv/config/config-dev.yaml"
	} else {
		configFilePath = "goods-srv/config/config-prod.yaml"
	}

	initialize.BindConfig(configFilePath, global.ServerConfig)
}
