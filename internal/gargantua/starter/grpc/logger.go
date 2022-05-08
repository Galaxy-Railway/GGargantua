package grpc

import "go.uber.org/zap"

var (
	logger *zap.SugaredLogger
)

func InitLogger(l *zap.SugaredLogger) {
	logger = l.Named("infra | interceptor")
}

func Logger() *zap.SugaredLogger {
	return logger
}
