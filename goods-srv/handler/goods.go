package handler

import (
	goodsProto "github.com/shengshunyan/mxshop-proto/goods/proto"
)

type GoodsServer struct {
	goodsProto.UnimplementedGoodsServer
}
