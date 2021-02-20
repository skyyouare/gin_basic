package logger

import (
	"gin_basic/pkg/setting"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// error logger
var errorLogger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func init() {
	level := getLoggerLevel("debug")
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  setting.LogSetting.FileName,
		MaxSize:   setting.LogSetting.MaxSize,
		LocalTime: setting.LogSetting.LocalTime,
		Compress:  setting.LogSetting.Compress,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	errorLogger = logger.Sugar()
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	errorLogger.Debug(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	errorLogger.Debugf(template, args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	errorLogger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	errorLogger.Infof(template, args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	errorLogger.Warn(args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	errorLogger.Warnf(template, args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	errorLogger.Error(args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	errorLogger.Errorf(template, args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	errorLogger.DPanic(args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	errorLogger.DPanicf(template, args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	errorLogger.Panic(args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	errorLogger.Panicf(template, args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	errorLogger.Fatal(args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	errorLogger.Fatalf(template, args...)
}
