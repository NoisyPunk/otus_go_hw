package internalgrpc

import (
	"context"
	"fmt"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/server/grpc/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type EventServer struct {
	server *grpc.Server
	port   string
	pb.UnimplementedEventsServer
}

func (e EventServer) CreateEvent(ctx context.Context, request *pb.CreateEventRequest) (*pb.EventResponse, error) {
	fmt.Println(request)
	return nil, nil
}

func (e EventServer) UpdateEvent(ctx context.Context, request *pb.EventActionRequest) (*pb.EventResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e EventServer) DeleteEvent(ctx context.Context, request *pb.EventActionRequest) (*pb.EventDeletionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e EventServer) DailyEventList(ctx context.Context, request *pb.IntervalListRequest) (*pb.EventList, error) {
	//TODO implement me
	panic("implement me")
}

func (e EventServer) WeeklyEventList(ctx context.Context, request *pb.IntervalListRequest) (*pb.EventList, error) {
	//TODO implement me
	panic("implement me")
}

func (e EventServer) MonthlyEventList(ctx context.Context, request *pb.IntervalListRequest) (*pb.EventList, error) {
	//TODO implement me
	panic("implement me")
}

func NewGRPCServer(port string) *EventServer {
	return &EventServer{
		server:                    grpc.NewServer(),
		port:                      port,
		UnimplementedEventsServer: pb.UnimplementedEventsServer{},
	}
}

func (e EventServer) Start(ctx context.Context) error {
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

func (e EventServer) Stop() {
	e.server.GracefulStop()
}
