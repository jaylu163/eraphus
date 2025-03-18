package config

import (
	"github.com/jaylu163/eraphus/internal/hades/logging"
	"github.com/jaylu163/eraphus/internal/manager"
)

type MagicAvatarConf struct {
	Queuelen     int `yaml:"Queuelen"`
	QueueWaitLen int `yaml:"QueueWaitLen"`
	CronTime     int `yaml:"CronTime"`
}

func Init() {

	// init log
	logging.LogInit()

	// init http client
	manager.NewRestCli()

	/*	// consul register
		consulConf := consul.DiscoveryConfig{
			ID:      "9527",
			Name:    config.GetHadesConf().ServerConf.ServerName,
			Tags:    []string{"a", "b"},
			Port:    config.GetHadesConf().ServerConf.Port,
			Address: "localhost", //通过ifconfig查看本机的eth0的ipv4地址
		}
		consul.Register(consulConf)
	*/
}
