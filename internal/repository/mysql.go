package repository

import (
	"context"
	"fmt"

	"github.com/jaylu163/eraphus/internal/hades/config"
	"github.com/jaylu163/eraphus/internal/hades/engine"
	"github.com/jaylu163/eraphus/internal/hades/logging"
	"github.com/jaylu163/eraphus/internal/models"
)

func InitMysql() {

	// 配置中的服务资源都在这块初始化
	err := config.InitMysql("weishiji")
	if err != nil {
		logging.Errorf("init mysql weishiji err:%v", err)
	}
	fmt.Println("err:", err)
	/*err = config.InitMysql("user")
	if err != nil {
		logs.Errorf("init mysql user err:%v", err)
	}*/
}

func GetMysqlConn(name string) *engine.MysqlConn {
	conn := engine.GetMysqlConn(name)
	return conn
}

func GetListById(ctx context.Context, id int64) (*models.TVPlayInfo, error) {
	info := &models.TVPlayInfo{}
	err := GetMysqlConn("weishiji").DB.First(info).Error

	return info, err
}
