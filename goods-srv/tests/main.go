package main

import (
	"context"
	"fmt"
	goodsProto "github.com/shengshunyan/mxshop-proto/goods/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var client goodsProto.GoodsClient

func main() {
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("new client failed" + err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic("close client failed" + err.Error())
		}
	}(conn)

	client = goodsProto.NewGoodsClient(conn)
	//testGetList()
	//testCreate()
	//testGetAllCategorysList()
	//testGetSubCategory()
	//testGoodsList()
	//testBatchGetGoods()
	testGetGoodsDetail()
}

func testGetList() {
	rsp, err := client.BrandList(context.Background(), &goodsProto.BrandFilterRequest{
		//Pn:    1,
		//PSize: 10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand)
	}
}

func testCreate() {
	brand, err := client.CreateBrand(context.Background(), &goodsProto.BrandRequest{
		Name: "娃哈哈1",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(brand)
}

func testGetAllCategorysList() {
	rsp, err := client.GetAllCategorysList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.Total)
	for _, category := range rsp.Data {
		fmt.Println(category)
	}
}

func testGetSubCategory() {
	rsp, err := client.GetSubCategory(context.Background(), &goodsProto.CategoryListRequest{Id: 135475})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.Total)
}

func testGoodsList() {
	rsp, err := client.GoodsList(context.Background(), &goodsProto.GoodsFilterRequest{
		KeyWords: "苹果",
		//IsHot:    true,
		TopCategory: 130358,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.Total)
}

func testBatchGetGoods() {
	rsp, err := client.BatchGetGoods(context.Background(), &goodsProto.BatchGoodsIdInfo{
		Id: []int32{421, 422},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.Total)
}

func testGetGoodsDetail() {
	rsp, err := client.GetGoodsDetail(context.Background(), &goodsProto.GoodInfoRequest{
		Id: 421,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}
