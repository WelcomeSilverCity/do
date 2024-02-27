package do

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/WelcomeSilverCity/do/global"
)

func InitConfig(filePath string) {
	v := viper.New()
	v.SetConfigFile(filePath)
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

func ReadNacos() {
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: global.AllConfig.Nacos.Host,
			Port:   uint64(global.AllConfig.Nacos.Port),
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         global.AllConfig.Nacos.NamespaceId,
		NotLoadCacheAtStart: true,
		LogDir:              global.AllConfig.Nacos.LogDir,
		CacheDir:            global.AllConfig.Nacos.CacheDir,
		LogLevel:            global.AllConfig.Nacos.LogLevel,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
		return
	}

	//获取读取的
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "demo_test.json",
		Group:  "dev"})

	json.Unmarshal([]byte(content), &global.AllConfig)

	configClient.ListenConfig(vo.ConfigParam{
		DataId: fmt.Sprintf("%s", global.AllConfig.Nacos.DataId),
		Group:  fmt.Sprintf("%s", global.AllConfig.Nacos.Group),
		OnChange: func(namespace, group, dataId, data string) {
			err = json.Unmarshal([]byte(data), &global.AllConfig)
			if err != nil {
				zap.S().Panic(err)
			}
			zap.S().Debugf("最新内容：%s", global.AllConfig)
		},
	})
}
