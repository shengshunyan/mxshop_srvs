package initialize

import (
	"bytes"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"mxshop_srvs/common/config"
	"mxshop_srvs/common/utils"
)

func BindConfig(path string, serverConfig *config.ServerConfig) {
	// 本地读取nacos信息
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	nacosConfig := config.NacosConfig{}
	if err := v.Unmarshal(&nacosConfig); err != nil {
		panic(err)
	}
	zap.S().Infow("[config] get nacos info", "nacosConfig", &nacosConfig)

	// 远程读取nacos配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: nacosConfig.Host,
			Port:   nacosConfig.Port,
		},
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		Group:  nacosConfig.Group,
		DataId: nacosConfig.DataId,
	})
	parseNacosConfig(content, serverConfig)
	// 动态监听
	err = configClient.ListenConfig(vo.ConfigParam{
		Group:  nacosConfig.Group,
		DataId: nacosConfig.DataId,
		OnChange: func(namespace, group, dataId, data string) {
			parseNacosConfig(content, serverConfig)
		},
	})
	if err != nil {
		panic(err)
	}

	// 动态监控功能
	//v.WatchConfig()
	//v.OnConfigChange(func(e fsnotify.Event) {
	//	if err := v.ReadInConfig(); err != nil {
	//		panic(err)
	//	}
	//	if err := v.Unmarshal(&global.ServerConfig); err != nil {
	//		panic(err)
	//	}
	//	zap.S().Infow("[config] watch config change", "serverConfig", &global.ServerConfig)
	//})
}

// 使用 Viper 解析 YAML 内容
func parseNacosConfig(content string, container *config.ServerConfig) {
	yamlViper := viper.New()
	yamlViper.SetConfigType("yaml")
	err := yamlViper.ReadConfig(bytes.NewBuffer([]byte(content)))
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
	err = yamlViper.Unmarshal(&container)
	if err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	// 设置动态端口
	port, err := utils.GetFreePort()
	if err == nil {
		container.Port = port
	}
}
