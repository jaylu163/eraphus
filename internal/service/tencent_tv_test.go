package service

import (
	"fmt"
	selfConfig "github.com/jaylu163/eraphus/config"
	"github.com/jaylu163/eraphus/manager"
	"golang.org/x/net/context"
	"os"
	"strings"
	"testing"

	"github.com/jaylu163/eraphus/internal/hades/config"
)

func TestGetVideoList(t *testing.T) {
	suffix := os.Getenv("env_scheduler_suffix")

	filePath := "../env/config.yaml"
	if suffix != "" {
		filePath = strings.Join([]string{"../env/config.yaml", suffix}, ".")
	}
	fmt.Println(filePath, suffix)
	err := config.ConfigInit(filePath)
	if err != nil {
		fmt.Println("load config file err:", err)
		return
	}

	// init
	selfConfig.Init()
	manager.NewRestCli()

	GetHotRec(context.TODO(), "https://v.qq.com/channel/tv")
	list, err := GetVideoList(context.Background(), "https://v.qq.com/x/cover/mzc00200x1fpzo9.html")
	if err != nil {
		t.Errorf("GetVideoList err:%v", err)
		return
	}
	t.Logf("list:%v", list)

}
