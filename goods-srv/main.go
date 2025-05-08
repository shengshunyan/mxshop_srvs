package main

import (
	"fmt"
	goodsProto "github.com/shengshunyan/mxshop-proto/goods/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	cInitialize "mxshop_srvs/common/initialize"
	"mxshop_srvs/goods-srv/global"
	"mxshop_srvs/goods-srv/handler"
	"mxshop_srvs/goods-srv/initialize"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 初始化logger
	cInitialize.InitLogger()
	defer cInitialize.CloseLogger()

	// 初始化config
	initialize.InitConfig()

	// 初始化数据库连接
	initialize.InitDB()
	defer initialize.CloseDB()

	server := grpc.NewServer()
	// 注册服务
	goodsProto.RegisterGoodsServer(server, &handler.GoodsServer{})
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

	zap.S().Infow("run server success", "host", serverConfig.Host, "port", serverConfig.Port)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("close server")
}
