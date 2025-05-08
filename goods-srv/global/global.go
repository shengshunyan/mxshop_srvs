package global

import (
	"gorm.io/gorm"
	"mxshop_srvs/common/config"
)

var (
	DB           *gorm.DB
	ServerConfig = &config.ServerConfig{}
)
