package multiple_steps

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func InjectLogger(globalLogger *zap.SugaredLogger) {
	logger = globalLogger.Named("app | multiple_step")
}

func Logger() *zap.SugaredLogger {
	return logger
}
