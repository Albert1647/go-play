package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	// ts -> timestamp
	config.EncoderConfig.TimeKey = "timestamp"
	// Time Epoch -> ISO8601
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// Disable stacktrace
	// config.EncoderConfig.StacktraceKey = ""
	var err error
	// Skip caller to skip this file call
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

// Shorten logs call
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}
