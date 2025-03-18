package consul

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"

	consulapi "github.com/hashicorp/consul/api"
)

type DiscoveryConfig struct {
	ID      string
	Name    string
	Tags    []string
	Port    int
	Address string
}

var consulAddress = "127.0.0.1:8500"

func Register(dis DiscoveryConfig) error {
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Printf("create consul client : %v\n", err.Error())
	}
	registration := &consulapi.AgentServiceRegistration{
		ID:      dis.ID,
		Name:    dis.Name,
		Port:    dis.Port,
		Tags:    dis.Tags,
		Address: dis.Address,
	}
	// 启动tcp的健康检测，注意address不能使用127.0.0.1或者localhost，因为consul-agent在docker容器里，如果用这个的话，
	// consul会访问容器里的port就会出错，一直检查不到实例
	check := &consulapi.AgentServiceCheck{}
	check.TCP = fmt.Sprintf("%s:%d", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "6s"
	registration.Check = check

	if err := client.Agent().ServiceRegister(registration); err != nil {
		fmt.Printf("register to consul error: %v\n", err.Error())
		return err
	}
	return nil
}

func Deregister(serviceId string) {
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Printf("create consul client : %v\n", err.Error())
	}
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		if err := client.Agent().ServiceDeregister(serviceId); err != nil {
			zap.S().Infof("注销user-service[%s] error", serviceId)
			os.Exit(1)
		}
		zap.S().Infof("注销user-service[%s] success", serviceId)
		os.Exit(0)
	}()
}
