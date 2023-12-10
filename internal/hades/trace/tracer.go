package trace

import (
	"context"
	"github.com/google/uuid"
)

// 定义键名
const (
	TraceIDKey      = "trace_id"
	UserIDKey       = "user_id"
	SpanTitleKey    = "span_title"
	SpanFunctionKey = "span_function"
	UserNameKey     = "user_name"
	TagKey          = "tag"
	VersionKey      = "version"
	StackKey        = "stack"
)

type (
	traceIDKey  struct{}
	userIDKey   struct{}
	userNameKey struct{}
	tagKey      struct{}
	stackKey    struct{}
)

type SpanOptions struct {
	Title    string
	FuncName string
}

func GenerateTraceId() string {
	return uuid.New().String()
}

// NewTraceIDContext 创建跟踪ID上下文
func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

// FromTraceIDContext 从上下文中获取跟踪ID
func FromTraceIDContext(ctx context.Context) string {
	v := ctx.Value(traceIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewUserIDContext 创建用户ID上下文
func NewUserIDContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

// FromUserIDContext 从上下文中获取用户ID
func FromUserIDContext(ctx context.Context) string {
	v := ctx.Value(userIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func NewUserNameContext(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, userNameKey{}, userName)
}

func FromUserNameContext(ctx context.Context) string {
	v := ctx.Value(userNameKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagKey{}, tag)
}

func FromTagContext(ctx context.Context) string {
	v := ctx.Value(tagKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func NewStackContext(ctx context.Context, stack error) context.Context {
	return context.WithValue(ctx, stackKey{}, stack)
}

func FromStackContext(ctx context.Context) error {
	v := ctx.Value(stackKey{})
	if v != nil {
		if s, ok := v.(error); ok {
			return s
		}
	}
	return nil
}

// SpanOption 定义跟踪单元的数据项
type SpanOption func(*SpanOptions)

// SetSpanTitle 设置跟踪单元的标题
func SetSpanTitle(title string) SpanOption {
	return func(o *SpanOptions) {
		o.Title = title
	}
}

// SetSpanFuncName 设置跟踪单元的函数名
func SetSpanFuncName(funcName string) SpanOption {
	return func(o *SpanOptions) {
		o.FuncName = funcName
	}
}
