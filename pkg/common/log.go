package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	once   = &sync.Once{}
	logger *zap.Logger
)

func InitLogger(config LogConfig) (err error) {
	once.Do(func() {
		var encoderCfg zapcore.EncoderConfig
		if config.EnvType == ProductEnv {
			encoderCfg = zap.NewProductionEncoderConfig()
		} else if config.EnvType == DevelopEnv {
			encoderCfg = zap.NewDevelopmentEncoderConfig()
		}
		encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.00")
		atom := zap.NewAtomicLevel()
		logger = zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			atom,
		))
	})
	logger.WithOptions()
	return
}

func Logger() *zap.SugaredLogger {
	return logger.Sugar()
}
