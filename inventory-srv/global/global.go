package global

import (
	"gorm.io/gorm"
	"mxshop_srvs/inventory-srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig = &config.ServerConfig{}
)
