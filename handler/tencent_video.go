package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jaylu163/eraphus/service"
	"net/http"
)

func GetRec(ctx *gin.Context) {
	hotList, err := service.GetHotRec(ctx, "https://v.qq.com/x/cover/mcv8hkc8zk8lnov/y0047i8dkgc.html?start=973&cut_vid=u0047ur59xr&scene_id=3")
	ctx.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"data":     hotList,
		"err_msg":  err,
	})
}

func GetVideoList(ctx *gin.Context) {
	c := ctx.Request.Context()
	hotList, err := service.GetVideoList(c, "https://v.qq.com/x/cover/7q544xyrava3vxf/p0033468jnx.html")
	ctx.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"data":     hotList,
		"err_msg":  err,
	})
}
