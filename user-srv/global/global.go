package global

import (
	"gorm.io/gorm"
	"mxshop_srvs/user-srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig = &config.ServerConfig{}
)
