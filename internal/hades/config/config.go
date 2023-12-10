package config

import "github.com/jaylu163/eraphus/internal/hades/logs"

// HadesConfig 基础框架配置共用参数结构体
/*
 mysql yaml配置
 redis yaml配置
 kafka yaml配置
 elasticsearch yaml配置
 etcd
 mongo
 ... etc
*/

// HadesConfig
type HadesConfig struct {
	LogConf       logs.LogConf        `yaml:"LogConf" toml:"log_conf"`
	ServerConf    ServerConf          `yaml:"ServerConf" toml:"server_conf"`
	ClientConf    []ClientConf        `yaml:"ClientConf" toml:"client_conf"`
	Mysql         []MysqlConf         `yaml:"Mysql" toml:"mysql"`
	Redis         []RedisConf         `yaml:"Redis" toml:"redis"`
	Kafka         []KafkaConf         `yaml:"Kafka" toml:"kafka"`
	Mongo         []MongoConf         `yaml:"Mongo" toml:"mongo"`
	Elasticsearch []ElasticsearchConf `yaml:"Elasticsearch" toml:"elasticsearch"`
	Rabbitmq      []RabbitmqConf      `yaml:"Rabbitmq" toml:"rabbitmq"`
}

type ServerConf struct {
	ServerName string `yaml:"ServerName" toml:"server_name"` // ServerName 使用三段.分割
	Port       int    `yaml:"Port" toml:"port"`
}
type ClientConf struct {
	ClientName string `yaml:"ClientName" toml:"client_name"`
	Endpoint   string `yaml:"Endpoint" toml:"endpoint"`
	Protocol   string `yaml:"Protocol" toml:"protocol"`
	Timeout    int    `yaml:"Timeout" toml:"timeout"` // 服务超时时间
	RetryTimes int    `yaml:"RetryTimes" toml:"retry_times"`
}

type MysqlConf struct {
	Name     string `yaml:"Name" toml:"name"` //
	Host     string `yaml:"Host" toml:"host"`
	Password string `yaml:"Password" toml:"password"`
	Username string `yaml:"Username" toml:"username"`
	Port     int    `yaml:"Port" toml:"port"`
	DBName   string `yaml:"DBName" toml:"dbname"`
}

type RedisConf struct {
	Name     string `yaml:"Name" toml:"name"`
	Host     string `yaml:"Host" toml:"host"`
	Port     int    `yaml:"Port" toml:"port"`
	DB       int    `yaml:"DB" toml:"db"`
	Password string `yaml:"Password" toml:"password"`
}

type KafkaConf struct {
}

type MongoConf struct {
}

type ElasticsearchConf struct {
}

type RabbitmqConf struct {
}
