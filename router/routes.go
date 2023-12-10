package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jaylu163/eraphus/handler"
)

func InitRoute(route *gin.Engine) {
	route.Use(handler.RequestLogger(), gin.Recovery())
	tencentGroup := route.Group("/api/v1/tencent")

	tencentGroup.GET("/get_video_list", handler.GetVideoList)
}
