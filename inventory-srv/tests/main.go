package main

import (
	"context"
	"fmt"
	"github.com/shengshunyan/mxshop-proto/inventory/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
)

var invClient proto.InventoryClient

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

	invClient = proto.NewInventoryClient(conn)

	// 初始化库存
	//var i int32
	//for i = 421; i <= 840; i++ {
	//	_, err := invClient.SetInv(context.Background(), &proto.GoodsInvInfo{
	//		GoodsId: i,
	//		Num:     100,
	//	})
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//fmt.Println("设置库存成功")

	// 并发情况之下 库存无法正确的扣减
	var wg sync.WaitGroup
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go TestSell(&wg)
	}
	wg.Wait()
	fmt.Println("库存扣减成功")
}

func TestSell(wg *sync.WaitGroup) {
	/*
		1. 第一件扣减成功： 第二件： 1. 没有库存信息 2. 库存不足
		2. 两件都扣减成功
	*/
	defer wg.Done()
	_, err := invClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 421, Num: 1},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("库存-1")
}
