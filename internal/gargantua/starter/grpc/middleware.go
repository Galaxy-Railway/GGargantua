package grpc

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"time"
)

func SetGrpcInterceptor() []grpc.ServerOption {
	interceptor := make([]grpc.UnaryServerInterceptor, 0)
	interceptor = append(interceptor, LogInterceptor)
	interceptor = append(interceptor, RecoverInterceptor)

	options := make([]grpc.ServerOption, 0)
	options = append(options, grpc.ChainUnaryInterceptor(interceptor...))
	return options
}

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	remote, _ := peer.FromContext(ctx)
	remoteAddr := remote.Addr.String()

	in, _ := json.Marshal(req)
	inStr := string(in)
	Logger().Debug("ip", remoteAddr, "access_start", info.FullMethod, "in", inStr)
	Logger().Info("access_start", info.FullMethod)
	start := time.Now()
	defer func() {
		out, _ := json.Marshal(resp)
		outStr := string(out)
		duration := int64(time.Since(start) / time.Millisecond)
		if duration >= 500 {
			Logger().Debug("ip", remoteAddr, "access_end", info.FullMethod, "in", inStr, "out", outStr, "err", err, "duration/ms", duration)
		} else {
			Logger().Debug("ip", remoteAddr, "access_end", info.FullMethod, "in", inStr, "out", outStr, "err", err, "duration/ms", duration)
		}
		Logger().Info("access_end", info.FullMethod)
	}()

	resp, err = handler(ctx, req)

	return
}

func RecoverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	panicked := true

	defer func() {
		if r := recover(); r != nil || panicked {
			err = recoverFrom(ctx, r)
		}
	}()

	resp, err = handler(ctx, req)
	panicked = false
	return resp, err
}

func recoverFrom(_ context.Context, p interface{}) error {
	Logger().Errorf("got panic: %+v \n stack: %+v", p, errors.WithStack(errors.New("got panic")))
	return status.Errorf(codes.Internal, "%v", p)
}
