package main

import (
	"fmt"
	userProto "github.com/shengshunyan/mxshop-proto/user/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"mxshop_srvs/user-srv/global"
	"mxshop_srvs/user-srv/handler"
	"mxshop_srvs/user-srv/initialize"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 初始化logger
	initialize.InitLogger()
	defer initialize.CloseLogger()

	// 初始化config
	initialize.InitConfig()

	// 初始化数据库连接
	initialize.InitDB()

	server := grpc.NewServer()
	// 注册用户服务
	userProto.RegisterUserServer(server, &handler.UserServer{})
	// 注册健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 服务注册
	initialize.InitRegister()
	defer initialize.CloseRegister()

	serverConfig := global.ServerConfig
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	go func() {
		// 阻塞的方法
		err = server.Serve(lis)
		if err != nil {
			panic("failed to serve: " + err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("close server")
}
