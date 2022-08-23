package logging

import (
	"fmt"
	"io"
	"os"
	"time"

	"userInfoService/pkg/setting"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZapLogger *zap.Logger
var SugarLogger *zap.SugaredLogger

func init() {
	encoder := getEncoder()

	errorWriteSyncer := getErrorLogWriter()
	debugWriteSyncer := getDebugLogWriter()

	coreDebug := zapcore.NewCore(encoder, debugWriteSyncer, zapcore.DebugLevel)
	coreError := zapcore.NewCore(encoder, errorWriteSyncer, zapcore.ErrorLevel)

	core := zapcore.NewTee(coreError, coreDebug)
	// ZapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	// ZapLogger, _ = zap.NewProduction()
	ZapLogger := zap.New(core, zap.AddCaller())
	SugarLogger = ZapLogger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getErrorLogWriter() zapcore.WriteSyncer {

	errorLogPath := fmt.Sprintf(setting.Cfg.ErrorLogDir, time.Now().Format("2006"), time.Now().Format("01"), time.Now().Format("02"))

	file, err := os.OpenFile(errorLogPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		panic("日志文件打开失败")
	}

	ws := io.MultiWriter(file, os.Stdout)
	return zapcore.AddSync(ws)
}

func getDebugLogWriter() zapcore.WriteSyncer {

	errorLogPath := fmt.Sprintf(setting.Cfg.DebugLogDir, time.Now().Format("2006"), time.Now().Format("01"), time.Now().Format("02"))

	file, err := os.OpenFile(errorLogPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		panic("日志文件打开失败")
	}

	ws := io.MultiWriter(file, os.Stdout)
	return zapcore.AddSync(ws)
}
