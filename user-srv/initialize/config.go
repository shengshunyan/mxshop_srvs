package initialize

import (
	"mxshop_srvs/common/initialize"
	"mxshop_srvs/common/utils"
	"mxshop_srvs/user-srv/global"
)

func InitConfig() {
	env := utils.GetEnv()
	configFilePath := ""
	if env == "dev" {
		configFilePath = "user-srv/config/config-dev.yaml"
	} else {
		configFilePath = "user-srv/config/config-prod.yaml"
	}

	initialize.BindConfig(configFilePath, global.ServerConfig)
}
