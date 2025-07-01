package model

import (
	"mxshop_srvs/common/model"
)

type Inventory struct {
	model.BaseModel
	Goods   int32 `gorm:"type:int;index"`
	Stocks  int32 `gorm:"type:int"`
	Version int32 `gorm:"type:int"` //分布式锁的乐观锁
}
