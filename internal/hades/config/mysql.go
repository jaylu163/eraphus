package config

import (
	"fmt"

	"github.com/jaylu163/eraphus/internal/hades/engine"
	"github.com/jaylu163/eraphus/internal/hades/logs"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	loggerorm "gorm.io/gorm/logger"
)

var (
	mysqlDict = make(map[string]*mysqlProxy)
)

type mysqlProxy struct {
	host         string `yaml:"Host" toml:"host"`
	username     string `yaml:"Username" toml:"username"`
	port         int    `yaml:"Port" toml:"port"`
	password     string `yaml:"Password" toml:"password"`
	dbname       string `yaml:"DBName" toml:"dbname"`
	maxIdleConns int    `yaml:"MaxIdleConns" toml:"max_idle_conns"` //空闲连接池中连接的最大数量
	maxOpenConns int    `yaml:"MaxOpenConns" toml:"max_open_conns"` //打开数据库连接的最大数量
}

func mysqlMap(confList []MysqlConf) {
	for _, item := range confList {
		mysqlDict[item.Name] = &mysqlProxy{
			host:     item.Host,
			username: item.Username,
			password: item.Password,
			port:     item.Port,
			dbname:   item.DBName,
		}
	}
}

func InitMysql(name string) error {
	if myProxy, ok := mysqlDict[name]; ok {
		db := myProxy.new()
		err := engine.AddMysqlConn(name, db)
		return err
	}
	return nil
}

func (my *mysqlProxy) new() *engine.MysqlConn {
	mysqldsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(mysqldsn, my.username, my.password, my.host, my.port, my.dbname)
	fmt.Println("dsn:ssss:", dsn)
	newLogger := loggerorm.New(
		log.New(logs.LogInit(), "mysql:", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		loggerorm.Config{
			SlowThreshold:             time.Second,    // 慢 SQL 阈值
			LogLevel:                  loggerorm.Info, // 日志级别
			IgnoreRecordNotFoundError: false,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,           // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		fmt.Printf("mysql open dsn: %v  err: %v\n", dsn, err)
		return nil
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(my.maxIdleConns)
	sqlDB.SetMaxOpenConns(my.maxOpenConns)
	if err != nil {
		fmt.Printf("mysql init db.DB() dsn:%v err:%v", dsn, err)
		return nil
	}
	fmt.Printf("msyql conn dsn:%v success\n", dsn)
	conn := &engine.MysqlConn{
		DB: db,
	}
	return conn
}
