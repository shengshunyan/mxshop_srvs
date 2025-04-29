package config

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Namespace string `mapstructure:"namespace"`
	Group     string `mapstructure:"group"`
	DataId    string `mapstructure:"dataid"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name"`
	Host       string       `mapstructure:"host"`
	Port       int          `mapstructure:"port"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql"`
	ConsulInfo ConsulConfig `mapstructure:"consul"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"db"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
