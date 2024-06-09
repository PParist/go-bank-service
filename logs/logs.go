package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	var err error
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	//TODO: reference to caller and skip ... time
	Log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	Log.Info(message, fields...)
}
func Debug(message string, fields ...zap.Field) {
	Log.Debug(message, fields...)
}
func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		Log.Error(v.Error(), fields...)
	case string:
		Log.Error(v, fields...)
	}
}
