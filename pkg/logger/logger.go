package logger

import (
	"go.uber.org/zap"
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

func (l *Logger) SetupZapLogger(newLogger *zap.Logger) {
	l.mMutex.Lock()
	defer l.mMutex.Unlock()

	l.mZapLogger = newLogger
}

func (l *Logger) Errorf(f string, v ...interface{}) {
	l.mZapLogger.Error(f)
}

func (l *Logger) Warningf(f string, v ...interface{}) {
	l.mZapLogger.Warn(f)
}

func (l *Logger) Infof(f string, v ...interface{}) {
	l.mZapLogger.Info(f)
}

func (l *Logger) Debugf(f string, v ...interface{}) {
	l.mZapLogger.Debug(f)
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
	})
}
