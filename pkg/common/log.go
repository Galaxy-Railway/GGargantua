package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	GlobalLogger *zap.Logger
)

func NewLogger(pConfig *ProjectConfig) (logger *zap.Logger, err error) {
	config := pConfig.Log
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
	logger.WithOptions()
	GlobalLogger = logger
	return
}

//deprecated: this function should be delete, all logger should get by inject
func Logger() *zap.SugaredLogger {
	return GlobalLogger.Sugar()
}
