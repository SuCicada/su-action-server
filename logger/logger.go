package logger

import (
	"fmt"
	"reflect"

	"github.com/SuCicada/su-action-server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logClient *zap.Logger

func InitLog() {
	config := zap.NewDevelopmentConfig()
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig = encoderConfig

	var logLevel zapcore.Level
	switch utils.Get("LOG_LEVEL") {
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	default:
		logLevel = zap.DebugLevel
	}
	config.Level = zap.NewAtomicLevelAt(logLevel)
	logger, err := config.Build(zap.AddCallerSkip(2))
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logClient = logger
}

func getMsgStr(msgs ...interface{}) string {
	msgStr := ""
	for _, msg := range msgs {
		if msg == nil {
			msgStr = msgStr + fmt.Sprintf("[nil]")
			continue
		}
		if reflect.TypeOf(msg).String() == "string" {
			msgStr = msgStr + fmt.Sprintf("[%v]", msg)
		} else {
			msgStr = msgStr + fmt.Sprintf("[<%s>%v]", reflect.TypeOf(msg).String(), msg)
		}
	}
	return msgStr
}

func Debug(msgs ...interface{}) {
	logClient.Debug(getMsgStr(msgs))
}

func Info(msgs ...interface{}) {
	logClient.Info(getMsgStr(msgs))
}

func Warn(msgs ...interface{}) {
	logClient.Warn(getMsgStr(msgs))
}

func Error(msgs ...interface{}) {
	logClient.Error(getMsgStr(msgs))
}
