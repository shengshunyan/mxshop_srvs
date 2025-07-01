package initialize

import (
	"mxshop_srvs/common/initialize"
	"mxshop_srvs/common/utils"
	"mxshop_srvs/inventory-srv/global"
)

func InitConfig() {
	env := utils.GetEnv()
	configFilePath := ""
	if env == "dev" {
		configFilePath = "inventory-srv/config/config-dev.yaml"
	} else {
		configFilePath = "inventory-srv/config/config-prod.yaml"
	}

	initialize.BindConfig(configFilePath, global.ServerConfig)
}
