package config

type Config struct {
	Srvname string       `mapstructure:"srvname" json:"srvname"`
	Host    string       `mapstructure:"host" json:"host"`
	Port    int          `mapstructure:"port" json:"port"`
	Mysql   MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	Consul  ConsulConfig `mapstructure:"consul" json:"consul"`
	Nacos   NacosConfig  `mapstructure:"nacos" json:"nacos"`
	EsInfo  EsConfig     `mapstructure:"es" json:"es"`
}

type MysqlConfig struct {
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Ip   string `mapstructure:"ip" json:"ip"`
}

type EsConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	DataId      string `mapstructure:"data_id" json:"data_id"`
	Group       string `mapstructure:"group" json:"group"`
	NamespaceId string `mapstructure:"namespace_id" json:"namespace_id"`
	LogDir      string `mapstructure:"log_dir" json:"log_dir"`
	CacheDir    string `mapstructure:"cache_dir" json:"cache_dir"`
	LogLevel    string `mapstructure:"log_level" json:"log_level"`
}
