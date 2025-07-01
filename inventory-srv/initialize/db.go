package initialize

import (
	"mxshop_srvs/common/initialize"
	"mxshop_srvs/inventory-srv/global"
)

func InitDB() {
	mysqlConfig := global.ServerConfig.MysqlInfo
	global.DB = initialize.InitDB(mysqlConfig)
}

func CloseDB() {
	sqlDB, _ := global.DB.DB()
	_ = sqlDB.Close()
}
