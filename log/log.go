package log

import (
	"github.com/lisuiheng/fsm"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level type
type Level int8

// These are the different logging levels.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel = Level(zapcore.PanicLevel)
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel = Level(zapcore.FatalLevel)
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = Level(zapcore.ErrorLevel)
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel = Level(zapcore.WarnLevel)
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel = Level(zapcore.InfoLevel)
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel = Level(zapcore.DebugLevel)
)

// ParseLevel takes a string level and returns the Logrus log level constant.
func ParseLevel(lvl string) Level {
	switch lvl {
	case "panic":
		return PanicLevel
	case "fatal":
		return FatalLevel
	case "error":
		return ErrorLevel
	case "warn":
		return WarnLevel
	case "info":
		return InfoLevel
	default:
		return DebugLevel
	}
}

// LogrusLogger is a chronograf.Logger that uses logrus to process logs
type zapLogger struct {
	l *zap.SugaredLogger
}

func (ll *zapLogger) Debug(args ...interface{}) {
	ll.l.Debug(args...)
}

func (ll *zapLogger) Info(args ...interface{}) {
	ll.l.Info(args...)
}

func (ll *zapLogger) Warn(args ...interface{}) {
	ll.l.Warn(args...)
}

func (ll *zapLogger) Error(args ...interface{}) {
	ll.l.Error(args...)
}

func (ll *zapLogger) Fatal(args ...interface{}) {
	ll.l.Fatal(args...)
}

func (ll *zapLogger) Panic(args ...interface{}) {
	ll.l.Panic(args...)
}

func (ll *zapLogger) Debugf(template string, args ...interface{}) {
	ll.l.Debugf(template, args...)
}

func (ll *zapLogger) Infof(template string, args ...interface{}) {
	ll.l.Infof(template, args...)
}

func (ll *zapLogger) Warnf(template string, args ...interface{}) {
	ll.l.Warnf(template, args...)
}

func (ll *zapLogger) Errorf(template string, args ...interface{}) {
	ll.l.Errorf(template, args...)
}

func (ll *zapLogger) Fatalf(template string, args ...interface{}) {
	ll.l.Fatalf(template, args...)
}

func (ll *zapLogger) Panicf(template string, args ...interface{}) {
	ll.l.Panicf(template, args...)
}

// New wraps a logrus Logger
func New(level Level) fsm.Logger {
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.Level(level)),
		OutputPaths: []string{"stdout"},
	}.Build()
	sugar := logger.Sugar()
	return &zapLogger{
		l: sugar,
	}
}
