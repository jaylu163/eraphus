package logging

import (
	"context"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/jaylu163/eraphus/internal/hades/trace"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	//logger  *zap.SugaredLogger
	careLog *CareLog
)

const (
	DEBUG  = "DEBUG"
	INFO   = "INFO"
	WARN   = "WARN"
	ERROR  = "ERROR"
	DPANIC = "DPANIC"
	PANIC  = "PANIC"
	FATAL  = "FATAL"
)

type LogConf struct {
	LogPath    string `yaml:"LogPath" toml:"log_path" json:"log_path"` // 日志存储路径
	Level      string `yaml:"Level" toml:"level" toml:"level"`
	Prefix     string `yaml:"Prefix" toml:"prefix" json:"prefix"`
	MaxSize    int    `yaml:"MaxSize" toml:"max_size" json:"maxsize"`
	MaxAge     int    `yaml:"MaxAge" yaml:"max_age" json:"max_age"`
	MaxBackups int    `yaml:"MaxBackups" toml:"max_backups" json:"max_backups"`
}

type CareLog struct {
	logger     *zap.SugaredLogger
	prefix     string
	withFields WithFields
}

func LogInit() *CareLog {
	return careLog
}

// SugarInit 配置传入日志类并初始化
func SugarInit(conf *LogConf) {

	/*	logLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			lvl := zapcore.Level(0)
			if strings.ToUpper(conf.Level) == DEBUG {
				lvl = zapcore.DebugLevel
			}
			if strings.ToUpper(conf.Level) == INFO {
				lvl = zapcore.InfoLevel
			}
			if strings.ToUpper(conf.Level) == WARN {
				lvl = zapcore.WarnLevel
			}
			if strings.ToUpper(conf.Level) == ERROR {
				lvl = zapcore.ErrorLevel
			}
			if strings.ToUpper(conf.Level) == DPANIC {
				lvl = zapcore.DPanicLevel
			}
			if strings.ToUpper(conf.Level) == PANIC {
				lvl = zapcore.PanicLevel
			}
			return level >= lvl
		})
	*/
	//https://blog.51cto.com/u_15064650/4071962
	infoHook := getLogWriter(conf.LogPath+"/info.log", conf.MaxSize, conf.MaxBackups, conf.MaxAge)
	warnHook := getLogWriter(conf.LogPath+"/warn.log", conf.MaxSize, conf.MaxBackups, conf.MaxAge)
	errHook := getLogWriter(conf.LogPath+"/error.log", conf.MaxSize, conf.MaxBackups, conf.MaxAge)
	// 实现两个判断日志等级的interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})
	errLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	})
	newCore := zapcore.NewTee(
		zapcore.NewCore(getEncoder(), zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(getEncoder(), zapcore.AddSync(infoHook), infoLevel),
		zapcore.NewCore(getEncoder(), zapcore.AddSync(warnHook), warnLevel),
		zapcore.NewCore(getEncoder(), zapcore.AddSync(errHook), errLevel),
	)
	//newCore := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(infoHook, warnHook, errHook), logLevel)

	zapLogger := zap.New(newCore, zap.AddCaller(), zap.AddCallerSkip(0))
	zapLogger.Named(conf.Prefix) // 设置前缀

	careLog = &CareLog{
		logger: zapLogger.Sugar(),
		prefix: conf.Prefix,
	}
}

func getEncoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		MessageKey:          "msg",
		LevelKey:            "level",
		TimeKey:             "time",
		NameKey:             "logger",
		CallerKey:           "file",
		FunctionKey:         "func",
		StacktraceKey:       "stacktrace",
		SkipLineEnding:      false,
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         zapcore.CapitalColorLevelEncoder, //zapcore.CapitalLevelEncoder,
		EncodeTime:          zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration:      zapcore.StringDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		EncodeName:          zapcore.FullNameEncoder,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    " ", //日志间分隔符空格
	}
	//return zapcore.NewJSONEncoder(config) // json 格式输出
	return zapcore.NewConsoleEncoder(config)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	// 使用 lumberjack 归档切片日志
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  //文件名字
		MaxSize:    maxSize,   //最大存储大小 byte
		MaxBackups: maxBackup, //备份数量
		MaxAge:     maxAge,    //最大存储时间
		Compress:   true,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
	//https://blog.51cto.com/u_12855930/5785490
}

func stringToLogLevel(level string) zapcore.Level {
	switch level {
	case "fatal":
		return zap.FatalLevel
	case "error":
		return zap.ErrorLevel
	case "warn":
		return zap.WarnLevel
	case "warning":
		return zap.WarnLevel
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	}
	return zap.DebugLevel
}

func (log *CareLog) Write(p []byte) (n int, err error) {
	// parse
	// \d+.\d+[a-z]+   --->  latency
	//  /.*:\d+ 		--->  file
	//  [A-Z].*\n 		--->  syntax
	// \[rows:(\d+)\]  -----> records

	latency := regexp.MustCompile("\\d+.\\d+[a-z]+").Find(p)
	path := regexp.MustCompile(" /[A-Za-z].*:\\d+").Find(p)
	syntax := regexp.MustCompile(" [A-Z].*").Find(p)
	affectRows := regexp.MustCompile(`\[rows:(\d+)\]`).Find(p)
	_, fileName := filepath.Split(string(path))
	log.logger.
		With("trace_id", "").
		Infof("latency:%s file:%s sql:%s affect:%s",
			strings.Trim(string(latency), " "),
			strings.Trim(fileName, " "),
			strings.Trim(string(syntax), " "),
			strings.Trim(string(affectRows), " "),
		)
	return len(p), nil
}

type WithFields struct {
	Status int    `json:"status"`
	Method string `json:"method"`
	Path   string `json:"path"`
	Ip     string `json:"ip"`
	Cost   string `json:"cost"`
}

func (log *CareLog) WithField(field WithFields) {
	log.withFields = field
}

func (log *CareLog) WithFor(ctx context.Context, args ...interface{}) *zap.SugaredLogger {
	traceId := trace.FromTraceIDContext(ctx)
	return log.logger.Named(log.prefix).With("trace_id", traceId, "args", args)
}

func (log *CareLog) Infof(template string, args ...interface{}) {
	log.logger.Infof(template, args...)
}

func (log *CareLog) Warnf(template string, args ...interface{}) {
	log.logger.Warnf(template, args...)
}
func (log *CareLog) Errorf(template string, args ...interface{}) {
	log.logger.Errorf(template, args...)
}
func (log *CareLog) Fatalf(format string, args ...interface{}) {
	log.logger.Fatalf(format, args)
}

// LogMode 实现gorm 日志
func (log *CareLog) LogMode(logLevel zapcore.Level) interface{} {
	return log
}

// Info 实现gorm 日志
func (log *CareLog) Info(ctx context.Context, args ...interface{}) {
	log.logger.With("trace_id", trace.FromTraceIDContext(ctx)).Info(args)
}

// Warn 实现gorm 日志
func (log *CareLog) Warn(ctx context.Context, args ...interface{}) {
	log.logger.With("trace_id", trace.FromTraceIDContext(ctx)).Warn(args)
}

// Warn 实现gorm 日志
func (log *CareLog) Error(ctx context.Context, args ...interface{}) {
	log.logger.With("trace_id", trace.FromTraceIDContext(ctx)).Error(args)
}

// Trace 实现gorm 日志
func (log *CareLog) Trace(begin time.Time, fc func() (sql string, rowsAffected int64), err error) {

}

/*
Infow("Failed to fetch URL.",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
*/
// Infow key values ...
func (log *CareLog) Infow(msg string, keysAndValues ...interface{}) {
	log.logger.Infow(msg, keysAndValues)
}

func (log *CareLog) Errorw(msg string, keysAndValues ...interface{}) {
	log.logger.Errorw(msg, keysAndValues)
}

func WithField(field WithFields) {
	careLog.withFields = field
}

func WithFor(ctx context.Context, args ...interface{}) *zap.SugaredLogger {
	traceId := trace.FromTraceIDContext(ctx)

	fields := []interface{}{}
	fields = append(fields,
		"method", careLog.withFields.Method,
		"ip", careLog.withFields.Ip,
		"path", careLog.withFields.Path,
		"cost", careLog.withFields.Cost,
		"status", careLog.withFields.Status)
	return careLog.logger.With("trace_id", traceId, "args", args, "fields", fields)
}

func Infof(template string, args ...interface{}) {
	careLog.logger.Infof(template, args)
}

func Errorf(template string, args ...interface{}) {
	careLog.logger.Errorf(template, args)
}
