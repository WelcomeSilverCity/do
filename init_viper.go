package initialize

import (
	"fmt"
	"net"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"zg6-demo/shop_srv/goods_srv/global"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("D:\\zg6-demo\\shop_srv\\user_srv\\config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
		return
	}
	v.Unmarshal(&global.AllConfig)
	if global.AllConfig.Port == 0 {
		global.AllConfig.Port = GetRandTp()
	}
	fmt.Println(global.AllConfig)

	go func() {
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
			_ = v.ReadInConfig() // 读取配置数据
			_ = v.Unmarshal(&global.AllConfig)
			InitDB()
		})
	}()
}

func GetRandTp() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		zap.S().Panic(err)
		return 0
	}
	l, _ := net.ListenTCP("tcp", addr)
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}
