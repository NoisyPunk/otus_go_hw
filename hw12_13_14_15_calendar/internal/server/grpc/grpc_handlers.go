package internalgrpc

import (
	"context"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/server/grpc/pb"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	day   = "day"
	week  = "week"
	month = "month"
)

func (e *GRPCEventServer) CreateEvent(ctx context.Context,
	request *pb.CreateEventRequest,
) (*pb.EventResponse, error) {
	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}
	eventData := storage.Event{
		Title:        request.Title,
		DateAndTime:  request.DateAndTime.AsTime(),
		Duration:     int(request.Duration),
		Description:  request.Description,
		UserID:       userID,
		TimeToNotify: int(request.TimeToNotify),
	}
	event, err := e.application.CreateEvent(ctx, eventData, userID)
	if err != nil {
		return nil, err
	}

	response := &pb.EventResponse{
		EventId:      event.ID.String(),
		Title:        event.Title,
		DateAndTime:  timestamppb.New(event.DateAndTime),
		Duration:     int32(event.Duration),
		Description:  event.Description,
		UserId:       event.UserID.String(),
		TimeToNotify: int32(event.TimeToNotify),
	}
	return response, nil
}

func (e *GRPCEventServer) UpdateEvent(ctx context.Context,
	request *pb.EventUpdateRequest,
) (*pb.EventUpdateResponse, error) {
	eventID, err := uuid.Parse(request.EventId)
	if err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(request.Event.UserId)
	if err != nil {
		return nil, err
	}
	eventData := storage.Event{
		ID:           eventID,
		Title:        request.Event.Title,
		DateAndTime:  request.Event.DateAndTime.AsTime(),
		Duration:     int(request.Event.Duration),
		Description:  request.Event.Description,
		UserID:       userID,
		TimeToNotify: int(request.Event.TimeToNotify),
	}

	err = e.application.UpdateEvent(ctx, eventData.ID, eventData)
	if err != nil {
		return nil, err
	}

	response := &pb.EventUpdateResponse{
		EventId: eventData.ID.String(),
		Message: "event successfully updated",
	}
	return response, nil
}

func (e *GRPCEventServer) DeleteEvent(ctx context.Context,
	request *pb.EventDeletionRequest,
) (*pb.EventDeletionResponse, error) {
	eventID, err := uuid.Parse(request.EventId)
	if err != nil {
		return nil, err
	}

	err = e.application.DeleteEvent(ctx, eventID)
	if err != nil {
		return nil, err
	}

	response := &pb.EventDeletionResponse{
		EventId: eventID.String(),
		Message: "event successfully deleted",
	}
	return response, nil
}

func (e *GRPCEventServer) DailyEventList(ctx context.Context,
	request *pb.IntervalListRequest,
) (*pb.EventList, error) {
	return e.collectEventList(ctx, request, day)
}

func (e *GRPCEventServer) WeeklyEventList(ctx context.Context,
	request *pb.IntervalListRequest,
) (*pb.EventList, error) {
	return e.collectEventList(ctx, request, week)
}

func (e *GRPCEventServer) MonthlyEventList(ctx context.Context,
	request *pb.IntervalListRequest,
) (*pb.EventList, error) {
	return e.collectEventList(ctx, request, month)
}

func (e *GRPCEventServer) collectEventList(ctx context.Context,
	request *pb.IntervalListRequest, period string,
) (*pb.EventList, error) {
	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}
	dateTime := request.DateAndTime.AsTime()

	var events []storage.Event
	switch period {
	case day:
		events, err = e.application.EventsDailyList(ctx, dateTime, userID)
		if err != nil {
			return nil, err
		}
	case week:
		events, err = e.application.EventsWeeklyList(ctx, dateTime, userID)
		if err != nil {
			return nil, err
		}
	case month:
		events, err = e.application.EventsMonthlyList(ctx, dateTime, userID)
		if err != nil {
			return nil, err
		}
	}

	eventList := make([]*pb.EventResponse, 0)

	for _, event := range events {
		response := pb.EventResponse{
			EventId:      event.ID.String(),
			Title:        event.Title,
			DateAndTime:  timestamppb.New(event.DateAndTime),
			Duration:     int32(event.Duration),
			Description:  event.Description,
			UserId:       event.UserID.String(),
			TimeToNotify: int32(event.TimeToNotify),
		}
		eventList = append(eventList, &response)
	}
	return &pb.EventList{EventList: eventList}, nil
}
