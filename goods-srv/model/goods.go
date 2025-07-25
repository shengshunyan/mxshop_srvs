package model

import "mxshop_srvs/common/model"

type Category struct {
	model.BaseModel
	Name             string `gorm:"type:varchar(20);not null" json:"name"`
	Level            int32  `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool   `gorm:"type:boolean;not null;default:false" json:"is_tab"`
	ParentCategoryID int32
	ParentCategory   *Category
	SubCategory      []*Category `gorm:"foreignkey:ParentCategoryID;references:ID" json:"sub_category"`
}

func (Category) TableName() string {
	return "category"
}

type Brands struct {
	model.BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);not null;default:''"`
}

type GoodsCategoryBrand struct {
	model.BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand;unique;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;index:idx_category_brand;unique;not null"`
	Brands     Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

type Banner struct {
	model.BaseModel
	Image string `gorm:"type:varchar(200);not null;"`
	Url   string `gorm:"type:varchar(200);not null;"`
	Index int32  `gorm:"type:int;not null;default:1"`
}

type Goods struct {
	model.BaseModel

	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"`
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null"`
	ShipFree bool `gorm:"default:false;not null"`
	IsNew    bool `gorm:"default:false;not null"`
	IsHot    bool `gorm:"default:false;not null"`

	Name            string   `gorm:"type:varchar(50);not null"`
	GoodsSn         string   `gorm:"type:varchar(50);not null"`
	ClickNum        int32    `gorm:"type:int;default:0;not null"`
	SoldNum         int32    `gorm:"type:int;default:0;not null"`
	FavNum          int32    `gorm:"type:int;default:0;not null"`
	MarketPrice     float32  `gorm:"not null"`
	ShopPrice       float32  `gorm:"not null"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null"`
	Images          GormList `gorm:"type:varchar(1000);not null"`
	DescImages      GormList `gorm:"type:varchar(1000);not null"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null"`
}

func (Goods) TableName() string {
	return "goods"
}
