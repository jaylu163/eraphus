package manager

import (
	"github.com/go-resty/resty/v2"
	"time"
)

var restClient *resty.Client

func NewRestCli() {
	restClient = resty.New().SetTimeout(time.Second * 120)
}

func GetRestCli() *resty.Client {
	return restClient
}
