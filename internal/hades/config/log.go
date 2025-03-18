package config

import (
	"github.com/jaylu163/eraphus/internal/hades/logging"
)

func logInit(logConf *logging.LogConf) {
	logging.SugarInit(logConf)
}
