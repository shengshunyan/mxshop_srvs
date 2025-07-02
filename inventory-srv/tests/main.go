package main

import (
	"context"
	"fmt"
	"github.com/shengshunyan/mxshop-proto/inventory/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("new client failed" + err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic("close client failed" + err.Error())
		}
	}(conn)

	invClient := proto.NewInventoryClient(conn)
	var i int32
	for i = 421; i <= 840; i++ {
		_, err := invClient.SetInv(context.Background(), &proto.GoodsInvInfo{
			GoodsId: i,
			Num:     100,
		})
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("设置库存成功")
}
