package handler

import (
	"context"
	"encoding/json"
	"github.com/shengshunyan/mxshop-proto/goods/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"mxshop_srvs/common/utils/slice"
	"mxshop_srvs/goods-srv/global"
	"mxshop_srvs/goods-srv/model"
)

func (g *GoodsServer) GetAllCategorysList(context.Context, *emptypb.Empty) (*proto.CategoryListResponse, error) {
	var categorys []*model.Category
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	b, _ := json.Marshal(categorys)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

func (g *GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	resp := &proto.SubCategoryListResponse{}
	var category *model.Category
	result := global.DB.First(&category, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "分类不存在")
	}

	resp.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	var subCategories []*model.Category
	global.DB.Where(&model.Category{ParentCategoryID: category.ID}).Find(&subCategories)
	resp.SubCategorys = slice.Map(subCategories, func(category *model.Category) *proto.CategoryInfoResponse {
		return &proto.CategoryInfoResponse{
			Id:             category.ID,
			Name:           category.Name,
			Level:          category.Level,
			IsTab:          category.IsTab,
			ParentCategory: category.ParentCategoryID,
		}
	})
	return resp, nil
}

func (g *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	category := model.Category{
		Name:  req.Name,
		Level: req.Level,
		IsTab: req.IsTab,
	}
	if req.Level != 1 {
		//去查询父类目是否存在
		category.ParentCategoryID = req.ParentCategory
	}
	global.DB.Model(&model.Category{}).Create(category)
	return &proto.CategoryInfoResponse{Id: category.ID}, nil
}

func (g *GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	return &emptypb.Empty{}, nil
}

func (g *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	var category model.Category

	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryID = req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}

	global.DB.Save(&category)

	return &emptypb.Empty{}, nil
}
