package manager

import (
	"context"
	"github.com/jaylu163/eraphus/internal/hades/engine"
	"github.com/jaylu163/eraphus/internal/models"
)

func GetMysqlConn(name string) *engine.MysqlConn {
	conn := engine.GetMysqlConn(name)
	return conn
}

func GetListById(ctx context.Context, id int64) (*models.TVPlayInfo, error) {
	info := &models.TVPlayInfo{}
	err := GetMysqlConn("weishiji").DB.First(info).Error

	return info, err
}
