package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var (
	conf *HadesConfig
)

func LoadConf(filePath string) (*HadesConfig, error) {
	yamlBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	conf = &HadesConfig{}
	err = yaml.Unmarshal(yamlBytes, conf)
	return conf, err
}

func GetHadesConf() *HadesConfig {
	return conf
}

// ConfigInit 初始化配置
func ConfigInit(filePath string) error {
	conf, err := LoadConf(filePath)

	if err != nil {
		return err
	}
	if len(conf.Redis) > 0 {
		redisMap(conf.Redis)
	}
	if len(conf.Mysql) > 0 {
		mysqlMap(conf.Mysql)
	}

	// 日志初始化
	logInit(&conf.LogConf)
	// ....

	return nil
}
