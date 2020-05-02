package logger

import (
	"github.com/lni/dragonboat/v3/logger"
	"go.uber.org/zap"
	"sync"
)

var (
	loggerOnce sync.Once
	loggerIns  *Factory
)

type SugaredLogger struct {
	*zap.SugaredLogger
}

func (l *SugaredLogger) Warningf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func (l *SugaredLogger) SetLevel(logger.LogLevel) {

}

type Factory struct {
	mMutex       sync.Mutex
	mPackageName string
	mZapLogger   *zap.Logger
}

func (l *Factory) SetupZapLogger(newLogger *zap.Logger) {
	l.mMutex.Lock()
	defer l.mMutex.Unlock()

	l.mZapLogger = newLogger
}

func (l *Factory) GetZapLogger() *zap.Logger {
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

func GetLoggerFactory() *Factory {
	return loggerIns
}

func S(pkgName string) *SugaredLogger {
	return &SugaredLogger{loggerIns.GetZapLogger().Named(pkgName).Sugar()}
}

//func L() *zap.Logger {
//	return GetLoggerFactory().GetZapLogger()
//}

func L(pkgName string) *zap.Logger {
	return GetLoggerFactory().GetZapLogger().Named(pkgName)
}

func DragonboatLoggerFactory(pkgName string) logger.ILogger {
	return &SugaredLogger{loggerIns.GetZapLogger().Named(pkgName).Sugar()}
}

func init() {
	loggerOnce.Do(func() {
		l, e := zap.NewProduction()
		if e != nil {
			panic(e)
		}
		loggerIns = &Factory{mZapLogger: l}

		logger.SetLoggerFactory(DragonboatLoggerFactory)
	})
}
