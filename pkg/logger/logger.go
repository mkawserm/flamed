package logger

import (
	"fmt"
	"github.com/lni/dragonboat/v3/logger"
	"go.uber.org/zap"
	"strings"
	"sync"
)

var (
	loggerOnce sync.Once
	loggerIns  *Logger
)

type Logger struct {
	mMutex     sync.Mutex
	mZapLogger *zap.Logger
}

func (l *Logger) SetLevel(logger.LogLevel) {

}

func (l *Logger) SetupZapLogger(newLogger *zap.Logger) {
	l.mMutex.Lock()
	defer l.mMutex.Unlock()

	l.mZapLogger = newLogger
}

func (l *Logger) Errorf(f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	if strings.HasSuffix(msg, "\n") {
		msg = strings.TrimSuffix(msg, "\n")
	}
	l.mZapLogger.Error(msg)
}

func (l *Logger) Warningf(f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	if strings.HasSuffix(msg, "\n") {
		msg = strings.TrimSuffix(msg, "\n")
	}
	l.mZapLogger.Warn(msg)
}

func (l *Logger) Infof(f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	if strings.HasSuffix(msg, "\n") {
		msg = strings.TrimSuffix(msg, "\n")
	}
	l.mZapLogger.Info(msg)
}

func (l *Logger) Debugf(f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	if strings.HasSuffix(msg, "\n") {
		msg = strings.TrimSuffix(msg, "\n")
	}
	l.mZapLogger.Debug(msg)
}

func (l *Logger) Panicf(f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	if strings.HasSuffix(msg, "\n") {
		msg = strings.TrimSuffix(msg, "\n")
	}
	l.mZapLogger.Panic(msg)
}

func (l *Logger) GetZapLogger() *zap.Logger {
	l.mMutex.Lock()
	defer l.mMutex.Unlock()

	if l.mZapLogger == nil {
		log, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		l.mZapLogger = log
	}

	return l.mZapLogger
}

func GetLogger() *Logger {
	return loggerIns
}

func GetZapLogger() *zap.Logger {
	return GetLogger().GetZapLogger()
}

func init() {
	loggerOnce.Do(func() {
		l, e := zap.NewProduction()
		if e != nil {
			panic(e)
		}
		loggerIns = &Logger{mZapLogger: l}

		logger.SetLoggerFactory(func(pkgName string) logger.ILogger {
			return loggerIns
		})
	})
}
