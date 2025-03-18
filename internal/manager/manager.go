package manager

import (
	"fmt"
	"github.com/jaylu163/eraphus/internal/hades/config"
	"github.com/jaylu163/eraphus/internal/hades/logging"
)

type Manager struct {
}

// 初始化项目
var ()

func Init() {

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
