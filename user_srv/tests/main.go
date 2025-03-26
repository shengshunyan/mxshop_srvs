package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mxshop_srvs/user_srv/proto"
)

var client proto.UserClient

func main() {
	conn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("new client failed" + err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic("close client failed" + err.Error())
		}
	}(conn)

	client = proto.NewUserClient(conn)
	//testCreateUser()
	testGetUserList()
}

func testGetUserList() {
	rsp, err := client.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.Total)
	for _, user := range rsp.Data {
		fmt.Println(user)
		check, err2 := client.CheckPassword(context.Background(), &proto.CheckInfo{
			Password:          "123456",
			EncryptedPassword: user.Password,
		})
		if err2 != nil {
			panic(err2)
		}
		fmt.Println("check result", check)
	}
}

func testCreateUser() {
	user, err := client.CreateUser(context.Background(), &proto.CreateUserInfo{
		Password: "123456",
		Nickname: "Peng",
		Mobile:   "136917",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
