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

	mLoggers        map[string]*zap.Logger
	mSugaredLoggers map[string]*zap.SugaredLogger
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

	if l.mLoggers == nil {
		l.mLoggers = make(map[string]*zap.Logger)
	}

	if l.mSugaredLoggers == nil {
		l.mSugaredLoggers = make(map[string]*zap.SugaredLogger)
	}

	return l.mZapLogger
}

func (l *Factory) S(pkgName string) *zap.SugaredLogger {
	l.mMutex.Lock()
	defer l.mMutex.Unlock()

	v, ok := l.mSugaredLoggers[pkgName]
	if ok {
		return v
	} else {
		sl := l.mZapLogger.Named(pkgName).Sugar()
		l.mSugaredLoggers[pkgName] = sl
		return sl
	}
}

func (l *Factory) L(pkgName string) *zap.Logger {
	l.mMutex.Lock()
	defer l.mMutex.Unlock()

	v, ok := l.mLoggers[pkgName]
	if ok {
		return v
	} else {
		nl := l.mZapLogger.Named(pkgName)
		l.mLoggers[pkgName] = nl
		return nl
	}
}

func GetLoggerFactory() *Factory {
	return loggerIns
}

func CS(pkgName string) *SugaredLogger {
	return &SugaredLogger{loggerIns.S(pkgName)}
}

// L returns zap logger
func L(pkgName string) *zap.Logger {
	return loggerIns.L(pkgName)
}

// S returns zap sugared logger
func S(pkgName string) *zap.SugaredLogger {
	return loggerIns.S(pkgName)
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
		loggerIns = &Factory{
			mZapLogger:      l,
			mLoggers:        make(map[string]*zap.Logger),
			mSugaredLoggers: make(map[string]*zap.SugaredLogger),
		}

		logger.SetLoggerFactory(DragonboatLoggerFactory)
	})
}
