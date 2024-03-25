package internalgrpc

import (
	"context"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/server/grpc/pb"
	"google.golang.org/grpc"
	"net"
)

type EventServer struct {
	server *grpc.Server
	pb.UnimplementedEventsServer
}

func (e EventServer) CreateEvent(ctx context.Context, request *pb.CreateEventRequest) (*pb.EventResponse, error) {
	//TODO implement me
	panic("implement me")
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

func NewGRPCServer() *EventServer {
	return new(EventServer)
}

func Run(port string) error {
	eventServer := NewGRPCServer()
	eventServer.server = grpc.NewServer()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	pb.RegisterEventsServer(eventServer.server, eventServer)

	return eventServer.server.Serve(listener)
}
