package internalgrpc

import (
	"context"
	"time"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func (e *GRPCEventServer) loggingInterceptor(ctx context.Context, request interface{},
	serverInfo *grpc.UnaryServerInfo, grpcHandler grpc.UnaryHandler,
) (interface{}, error) {
	l := logger.FromContext(e.ctx)
	start := time.Now()
	dateTime := time.DateTime

	response, err := grpcHandler(ctx, request)

	l.Info(
		"grpc request stats",
		zap.Duration("latency", time.Since(start)),
		zap.String("date_time", dateTime),
		zap.String("method", serverInfo.FullMethod),
	)
	return response, err
}
