package config

import (
	"github.com/jaylu163/eraphus/internal/hades/logs"
	"github.com/jaylu163/eraphus/manager"
)

type MagicAvatarConf struct {
	Queuelen     int `yaml:"Queuelen"`
	QueueWaitLen int `yaml:"QueueWaitLen"`
	CronTime     int `yaml:"CronTime"`
}

func Init() {

	// init http client
	manager.NewRestCli()

	// init log
	logs.LogInit()
}
