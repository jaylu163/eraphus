package config

import (
	"github.com/jaylu163/eraphus/internal/hades/logs"
)

func logInit(logConf *logs.LogConf) {
	logs.SugarInit(logConf)
}
