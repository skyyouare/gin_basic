package logger

import (
	"gin_basic/pkg/setting"
	"time"

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

// Setup 初始化
func Setup() {
	level := getLoggerLevel("debug")
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  setting.LogSetting.FileName,
		MaxSize:   setting.LogSetting.MaxSize,
		LocalTime: setting.LogSetting.LocalTime,
		Compress:  setting.LogSetting.Compress,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.TimeKey = "time"
	encoder.FunctionKey = "func"
	encoder.StacktraceKey = "trace"
	encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
		enc.AppendString(t.In(cstSh).Format("2006-01-02 15:04:05"))
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
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

// Infow logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	errorLogger.Infow(msg, keysAndValues...)
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

// Errorw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	errorLogger.Errorw(msg, keysAndValues...)
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
