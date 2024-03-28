package internalgrpc

import (
	"context"
	"net"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/server/grpc/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCEventServer struct {
	application app.Application
	server      *grpc.Server
	port        string
	pb.UnimplementedEventsServer
}

func (e GRPCEventServer) Start(ctx context.Context) error {
	l := logger.FromContext(ctx)

	listener, err := net.Listen("tcp", e.port)
	if err != nil {
		return err
	}
	pb.RegisterEventsServer(e.server, e)

	go func() error {
		err = e.server.Serve(listener)
		if err != nil {
			return err
		}
		return nil
	}()
	l.Debug("grpc server started", zap.String("server port", e.port))
	<-ctx.Done()
	return nil
}

func (e GRPCEventServer) Stop() {
	e.server.GracefulStop()
}

func NewGRPCServer(app app.Application, port string) *GRPCEventServer {
	return &GRPCEventServer{
		application:               app,
		server:                    grpc.NewServer(),
		port:                      port,
		UnimplementedEventsServer: pb.UnimplementedEventsServer{},
	}
}
