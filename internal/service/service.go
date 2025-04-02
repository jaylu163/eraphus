package service

import (
	"github.com/jaylu163/eraphus/internal/manager"
	"github.com/jaylu163/eraphus/internal/repository"
)

func Init() {
	// 服务初始化
	manager.Init()
	repository.InitRepository()
}
