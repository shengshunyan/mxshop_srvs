package handler

import (
	inventoryProto "github.com/shengshunyan/mxshop-proto/inventory/proto"
)

type InventoryServer struct {
	inventoryProto.UnimplementedInventoryServer
}
