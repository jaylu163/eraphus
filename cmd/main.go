package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	selfConfig "github.com/jaylu163/eraphus/internal/config"
	"github.com/jaylu163/eraphus/internal/hades/config"
	"github.com/jaylu163/eraphus/internal/hades/consul"
	"github.com/jaylu163/eraphus/internal/router"
	"github.com/jaylu163/eraphus/internal/service"
	"golang.org/x/net/context"

	"github.com/gin-gonic/gin"
	"github.com/jaylu163/eraphus/internal/backend"
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

	// 启动服务
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.GetHadesConf().ServerConf.Port),
		Handler: r,
	}
	go serverStart(srv)

	fmt.Println("server addr", srv.Addr)
	// 监听服务挂起
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			// 摘除consul
			consul.Deregister("9527")
			// 停止当前服务
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
			log.Println("eraphus.logic server exit now...")
			return
		case syscall.SIGHUP:
		default:
		}
	}

}

func serverStart(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
