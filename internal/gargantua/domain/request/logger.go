package request

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func InjectLogger(globalLogger *zap.SugaredLogger) {
	logger = globalLogger.Named("domain | request")
}

func Logger() *zap.SugaredLogger {
	return logger
}
