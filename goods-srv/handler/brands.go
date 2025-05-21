package handler

import (
	"context"
	"github.com/shengshunyan/mxshop-proto/goods/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"mxshop_srvs/common/utils"
	"mxshop_srvs/common/utils/slice"
	"mxshop_srvs/goods-srv/global"
	"mxshop_srvs/goods-srv/model"
)

func (g *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	resp := &proto.BrandListResponse{}

	var count int64
	global.DB.Model(&model.Brands{}).Count(&count)
	resp.Total = int32(count)

	var brands []*model.Brands
	resulut := global.DB.Scopes(utils.Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if resulut.Error != nil {
		return nil, resulut.Error
	}

	resp.Data = slice.Map(brands, func(source *model.Brands) *proto.BrandInfoResponse {
		return &proto.BrandInfoResponse{
			Id:   source.ID,
			Name: source.Name,
			Logo: source.Logo,
		}
	})

	return resp, nil
}

func (g *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	if result := global.DB.Where(&model.Brands{Name: req.Name}).First(&model.Brands{}); result.RowsAffected == 1 {
		return nil, status.Error(codes.InvalidArgument, "品牌已存在")
	}

	brand := &model.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}
	if result := global.DB.Create(brand); result.Error != nil {
		return nil, result.Error
	}

	return &proto.BrandInfoResponse{Id: brand.ID}, nil
}

func (g *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Brands{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "品牌不存在")
	}

	return &emptypb.Empty{}, nil
}

func (g *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	if result := global.DB.First(&model.Brands{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "品牌不存在")
	}

	brand := &model.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}
	if result := global.DB.Updates(brand); result.Error != nil {
		return nil, result.Error
	}

	return &emptypb.Empty{}, nil
}
