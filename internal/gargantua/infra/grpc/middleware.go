package grpc

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"time"
)

var (
	logger *zap.SugaredLogger
)

func InitLogger(l *zap.SugaredLogger) {
	logger = l.Named("infra | interceptor")
}

func SetGrpcInterceptor() []grpc.ServerOption {
	interceptor := make([]grpc.UnaryServerInterceptor, 0)
	interceptor = append(interceptor, LogInterceptor)

	options := make([]grpc.ServerOption, 0)
	options = append(options, grpc.ChainUnaryInterceptor(interceptor...))
	return options
}

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	remote, _ := peer.FromContext(ctx)
	remoteAddr := remote.Addr.String()

	in, _ := json.Marshal(req)
	inStr := string(in)
	logger.Debug("ip", remoteAddr, "access_start", info.FullMethod, "in", inStr)
	logger.Info("access_start", info.FullMethod)
	start := time.Now()
	defer func() {
		out, _ := json.Marshal(resp)
		outStr := string(out)
		duration := int64(time.Since(start) / time.Millisecond)
		if duration >= 500 {
			logger.Debug("ip", remoteAddr, "access_end", info.FullMethod, "in", inStr, "out", outStr, "err", err, "duration/ms", duration)
		} else {
			logger.Debug("ip", remoteAddr, "access_end", info.FullMethod, "in", inStr, "out", outStr, "err", err, "duration/ms", duration)
		}
		logger.Info("access_end", info.FullMethod)
	}()

	resp, err = handler(ctx, req)

	return
}
