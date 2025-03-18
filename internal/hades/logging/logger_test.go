package logging

import (
	"context"
	"github.com/jaylu163/eraphus/internal/hades/trace"
	"testing"
)

func TestSugarInit(t *testing.T) {

	SugarInit(&LogConf{
		LogPath:    "logs_test",
		Level:      "info",
		Prefix:     "",
		MaxSize:    0,
		MaxAge:     0,
		MaxBackups: 0,
	})

	LogInit().WithFor(trace.NewTraceIDContext(context.Background(), trace.GenerateTraceId()), []string{"name", "zhang"}).
		Infof("info:%s %s", "aaa haha 芝士焗", "abc")

	careLog.WithFor(trace.NewTraceIDContext(context.Background(), trace.GenerateTraceId()), "func:", "context").Infof("你好:%s", "中国!😄")
}
