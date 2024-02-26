package initialize

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"

	"zg6-demo/shop_srv/goods_srv/global"
)

var Client *api.Client
var SrvId string

func init() {
	var err error
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	Client, err = api.NewClient(config)
	if err != nil {
		log.Println("Error", err)
	}
}

func InitConsul() {

	check := &api.AgentServiceCheck{
		GRPC:                           "10.2.178.95" + ":" + strconv.Itoa(global.AllConfig.Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	SrvId = fmt.Sprintf("%s", uuid.NewV4())
	message := &api.AgentServiceRegistration{
		ID:      SrvId,
		Name:    "goodsServer",
		Address: "127.0.0.1",
		Port:    global.AllConfig.Port,
		Check:   check,
	}
	err := Client.Agent().ServiceRegister(message)
	if err != nil {
		panic(err)
		return
	}
}

func End() {
	err := Client.Agent().ServiceDeregister(SrvId)
	if err != nil {
		panic(err)
		return
	}
}
