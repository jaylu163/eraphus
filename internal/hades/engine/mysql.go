package engine

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var (
	mysqlEngine map[string]*MysqlConn
)

type MysqlConn struct {
	DB *gorm.DB
}

func AddMysqlConn(name string, dbConn *MysqlConn) error {
	if mysqlEngine == nil {
		mysqlEngine = map[string]*MysqlConn{
			name: dbConn,
		}
		return nil
	}
	if _, ok := mysqlEngine[name]; ok {
		return errors.New(fmt.Sprintf("%s exists ,don't initialize repeated", name))
	}
	mysqlEngine[name] = dbConn
	return nil
}

func GetMysqlConn(name string) *MysqlConn {
	if mysqlEngine == nil {
		return nil
	}
	return mysqlEngine[name]
}

// 写一些方法
