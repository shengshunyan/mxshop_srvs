package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"mxshop_srvs/user-srv/global"
)

var client *api.Client
var serviceId = fmt.Sprintf("%s", uuid.NewV4())

// 服务注册
func InitRegister() {
	var err error

	serverConfig := global.ServerConfig
	consulConfig := serverConfig.ConsulInfo
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	client, err = api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	registration := &api.AgentServiceRegistration{
		Name:    serverConfig.Name,
		ID:      serviceId,
		Port:    serverConfig.Port,
		Address: serverConfig.Host,
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port),
			Timeout:                        "5s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "10s",
		},
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}

func CloseRegister() {
	err := client.Agent().ServiceDeregister(serviceId)
	if err != nil {
		zap.S().Errorf("ServiceDeregister err %v", err)
	}
}
