package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jaylu163/eraphus/internal/hades/logging"
	"github.com/jaylu163/eraphus/internal/hades/trace"
	"time"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values

		path := c.Request.URL.Path
		end := time.Now()
		latency := end.Sub(start)

		entry := logging.WithFields{
			Status: c.Writer.Status(),
			Method: c.Request.Method,
			Path:   path,
			Ip:     c.ClientIP(),
			Cost:   latency.String(),
			//"user-agent": c.Request.UserAgent(),
			//"time":       end.Format(time.DateTime),
		}

		traceId := c.Request.Header.Get("X-Request-ID")
		// 如果header头没有获取到trace ，生成一个trace
		if traceId == "" {
			traceId = trace.GenerateTraceId()
		}
		// 生成的traceId 重新传递给gin.Context
		c.Request = c.Request.WithContext(trace.NewTraceIDContext(c.Request.Context(), traceId))
		logging.WithField(entry)
		c.Next()
		// 服务请求的时候初始化 todo 待定，每次请求都初始化会浪费资源
		//service.Init()
	}
}
