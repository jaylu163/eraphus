package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jaylu163/eraphus/backend"
	selfConfig "github.com/jaylu163/eraphus/config"
	"github.com/jaylu163/eraphus/internal/hades/config"
	"github.com/jaylu163/eraphus/router"
	"github.com/jaylu163/eraphus/service"
	"os"
	"strconv"
	"strings"
)

func main() {

	//1.  load config
	suffix := os.Getenv("env_conf_suffix")
	filePath := "env/config.yaml"
	if suffix != "" {
		filePath = strings.Join([]string{"env/config.yaml", suffix}, ".")
	}
	fmt.Println(filePath, suffix)
	err := config.ConfigInit(filePath)
	if err != nil {
		fmt.Println("load config file err:", err)
		return
	}

	//2. init server
	r := gin.New()

	// init
	selfConfig.Init()

	// backend task
	backend.Start()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 服务初始化
	service.Init()
	// router New
	router.InitRoute(r)

	_ = r.Run(":" + strconv.Itoa(config.GetHadesConf().ServerConf.Port))
}
