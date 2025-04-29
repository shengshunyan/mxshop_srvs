package main

import (
	"context"
	"fmt"
	userProto "github.com/shengshunyan/mxshop-proto/user/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client userProto.UserClient

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

	client = userProto.NewUserClient(conn)
	testCreateUser()
	//testGetUserList()
}

func testGetUserList() {
	rsp, err := client.GetUserList(context.Background(), &userProto.PageInfo{
		Pn:    1,
		PSize: 10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.Total)
	for _, user := range rsp.Data {
		fmt.Println(user)
		check, err2 := client.CheckPassword(context.Background(), &userProto.CheckInfo{
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
	user, err := client.CreateUser(context.Background(), &userProto.CreateUserInfo{
		Password: "123456",
		Nickname: "xiaoming",
		Mobile:   "15754600156",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
