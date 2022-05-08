package job

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func InjectLogger(globalLogger *zap.SugaredLogger) {
	logger = globalLogger.Named("app | job")
}

func Logger() *zap.SugaredLogger {
	return logger
}
