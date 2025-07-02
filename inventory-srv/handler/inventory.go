package handler

import (
	"context"
	inventoryProto "github.com/shengshunyan/mxshop-proto/inventory/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"mxshop_srvs/inventory-srv/global"
	"mxshop_srvs/inventory-srv/model"
)

type InventoryServer struct {
	inventoryProto.UnimplementedInventoryServer
}

func (i *InventoryServer) SetInv(ctx context.Context, req *inventoryProto.GoodsInvInfo) (*emptypb.Empty, error) {
	var inv model.Inventory
	global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv)

	inv.Goods = req.GoodsId
	inv.Stocks = req.Num

	global.DB.Save(&inv)
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) InvDetail(ctx context.Context, req *inventoryProto.GoodsInvInfo) (*inventoryProto.GoodsInvInfo, error) {
	var inv model.Inventory
	if result := global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有库存信息")
	}
	return &inventoryProto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num:     inv.Stocks,
	}, nil
}

func (*InventoryServer) Sell(ctx context.Context, req *inventoryProto.SellInfo) (*emptypb.Empty, error) {
	//数据库基本的一个应用场景：数据库事务
	tx := global.DB.Begin()
	for _, goodInfo := range req.GoodsInfo {
		var inv model.Inventory

		if result := global.DB.Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
			tx.Rollback() //回滚之前的操作
			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
		}
		//判断库存是否充足
		if inv.Stocks < goodInfo.Num {
			tx.Rollback() //回滚之前的操作
			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
		}
		//扣减， 会出现数据不一致的问题 - 锁，分布式锁
		inv.Stocks -= goodInfo.Num
		tx.Save(&inv)
	}
	tx.Commit() // 需要自己手动提交操作
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) Reback(ctx context.Context, req *inventoryProto.SellInfo) (*emptypb.Empty, error) {
	//库存归还： 1：订单超时归还 2. 订单创建失败，归还之前扣减的库存 3. 手动归还
	tx := global.DB.Begin()
	for _, goodInfo := range req.GoodsInfo {
		var inv model.Inventory
		if result := global.DB.Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
			tx.Rollback() //回滚之前的操作
			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
		}

		//扣减， 会出现数据不一致的问题 - 锁，分布式锁
		inv.Stocks += goodInfo.Num
		tx.Save(&inv)
	}
	tx.Commit() // 需要自己手动提交操作
	return &emptypb.Empty{}, nil
}
