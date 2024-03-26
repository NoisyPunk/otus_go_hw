package internalgrpc

import (
	"context"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/server/grpc/pb"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (e GRPCEventServer) CreateEvent(ctx context.Context, request *pb.CreateEventRequest) (*pb.EventResponse, error) {
	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}
	eventData := storage.Event{
		Title:        request.Title,
		DateAndTime:  request.DateAndTime.AsTime(),
		Duration:     request.Duration.AsDuration(),
		Description:  request.Description,
		UserID:       userID,
		TimeToNotify: request.TimeToNotify.AsDuration(),
	}
	event, err := e.application.CreateEvent(ctx, eventData, userID)
	if err != nil {
		return nil, err
	}

	response := &pb.EventResponse{
		EventId:      event.ID.String(),
		Title:        event.Title,
		DateAndTime:  timestamppb.New(event.DateAndTime),
		Duration:     durationpb.New(event.Duration),
		Description:  event.Description,
		UserId:       event.UserID.String(),
		TimeToNotify: durationpb.New(event.TimeToNotify),
	}
	return response, nil
}

func (e GRPCEventServer) UpdateEvent(ctx context.Context, request *pb.EventActionRequest) (*pb.EventResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e GRPCEventServer) DeleteEvent(ctx context.Context, request *pb.EventActionRequest) (*pb.EventDeletionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e GRPCEventServer) DailyEventList(ctx context.Context, request *pb.IntervalListRequest) (*pb.EventList, error) {
	//TODO implement me
	panic("implement me")
}

func (e GRPCEventServer) WeeklyEventList(ctx context.Context, request *pb.IntervalListRequest) (*pb.EventList, error) {
	//TODO implement me
	panic("implement me")
}

func (e GRPCEventServer) MonthlyEventList(ctx context.Context, request *pb.IntervalListRequest) (*pb.EventList, error) {
	//TODO implement me
	panic("implement me")
}
