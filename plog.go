package plog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	//mode     string
	curLevel zapcore.Level
	sugar    *zap.SugaredLogger
)

func init() {
	InitPlogConfig("debug")
}

func InitPlogConfig(mode string) {
	initLogLevel(mode)
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.MessageKey = "message"
	cfg.LevelKey = "level"
	cfg.EncodeLevel = zapcore.LowercaseLevelEncoder

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), os.Stdout, curLevel)
	logger := zap.New(core)
	sugar = logger.Sugar().With(zap.String("Mode", mode))
	sugar.Named("plog")
}

func initLogLevel(mode string) {
	curLevel = zap.InfoLevel
	switch mode {
	case "release":
		curLevel = zap.InfoLevel
	case "debug":
		curLevel = zap.DebugLevel
	case "info":
		curLevel = zap.InfoLevel
	case "error":
		curLevel = zap.ErrorLevel
	case "panic":
		curLevel = zap.PanicLevel
	case "fatal":
		curLevel = zap.FatalLevel
	}
}

func Debugf(format string, v ...interface{}) {
	sugar.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	sugar.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	sugar.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	sugar.Errorf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	sugar.Panicf(format, v...)
}

func Close() {
	if sugar != nil {
		sugar.Infof("sugar sync!")
		sugar.Sync()
	}
}
