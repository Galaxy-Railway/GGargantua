package script

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func InjectLogger(globalLogger *zap.SugaredLogger) {
	logger = globalLogger.Named("domain | script")
}

func Logger() *zap.SugaredLogger {
	return logger
}
