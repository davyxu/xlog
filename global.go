package xlog

import "go.uber.org/zap"

var (
	slogger *zap.SugaredLogger
	logger  *zap.Logger
)

func Logger() *zap.Logger {
	return logger
}

func SetGlobal(l *zap.Logger) {
	logger = l
	slogger = l.Sugar()
}

func Debugf(format string, args ...interface{}) {
	slogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	slogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	slogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	slogger.Errorf(format, args...)
}

func Debugln(args ...interface{}) {
	slogger.Debug(args...)
}

func Infoln(args ...interface{}) {
	slogger.Info(args...)
}

func Warnln(args ...interface{}) {
	slogger.Warn(args...)
}

func Errorln(args ...interface{}) {
	slogger.Error(args...)
}

func init() {
	l := New(Config{})
	if l != nil {
		SetGlobal(l)
	}
}
