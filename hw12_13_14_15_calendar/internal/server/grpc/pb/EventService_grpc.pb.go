// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: EventService.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventsClient interface {
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	UpdateEvent(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventResponse, error)
	DeleteEvent(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventDeletionResponse, error)
	DailyEventList(ctx context.Context, in *IntervalListRequest, opts ...grpc.CallOption) (*EventList, error)
	WeeklyEventList(ctx context.Context, in *IntervalListRequest, opts ...grpc.CallOption) (*EventList, error)
	MonthlyEventList(ctx context.Context, in *IntervalListRequest, opts ...grpc.CallOption) (*EventList, error)
}

type eventsClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsClient(cc grpc.ClientConnInterface) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/event.Events/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) UpdateEvent(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/event.Events/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) DeleteEvent(ctx context.Context, in *EventActionRequest, opts ...grpc.CallOption) (*EventDeletionResponse, error) {
	out := new(EventDeletionResponse)
	err := c.cc.Invoke(ctx, "/event.Events/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) DailyEventList(ctx context.Context, in *IntervalListRequest, opts ...grpc.CallOption) (*EventList, error) {
	out := new(EventList)
	err := c.cc.Invoke(ctx, "/event.Events/DailyEventList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) WeeklyEventList(ctx context.Context, in *IntervalListRequest, opts ...grpc.CallOption) (*EventList, error) {
	out := new(EventList)
	err := c.cc.Invoke(ctx, "/event.Events/WeeklyEventList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) MonthlyEventList(ctx context.Context, in *IntervalListRequest, opts ...grpc.CallOption) (*EventList, error) {
	out := new(EventList)
	err := c.cc.Invoke(ctx, "/event.Events/MonthlyEventList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsServer is the server API for Events service.
// All implementations must embed UnimplementedEventsServer
// for forward compatibility
type EventsServer interface {
	CreateEvent(context.Context, *CreateEventRequest) (*EventResponse, error)
	UpdateEvent(context.Context, *EventActionRequest) (*EventResponse, error)
	DeleteEvent(context.Context, *EventActionRequest) (*EventDeletionResponse, error)
	DailyEventList(context.Context, *IntervalListRequest) (*EventList, error)
	WeeklyEventList(context.Context, *IntervalListRequest) (*EventList, error)
	MonthlyEventList(context.Context, *IntervalListRequest) (*EventList, error)
	mustEmbedUnimplementedEventsServer()
}

// UnimplementedEventsServer must be embedded to have forward compatible implementations.
type UnimplementedEventsServer struct {
}

func (UnimplementedEventsServer) CreateEvent(context.Context, *CreateEventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedEventsServer) UpdateEvent(context.Context, *EventActionRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedEventsServer) DeleteEvent(context.Context, *EventActionRequest) (*EventDeletionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedEventsServer) DailyEventList(context.Context, *IntervalListRequest) (*EventList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DailyEventList not implemented")
}
func (UnimplementedEventsServer) WeeklyEventList(context.Context, *IntervalListRequest) (*EventList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WeeklyEventList not implemented")
}
func (UnimplementedEventsServer) MonthlyEventList(context.Context, *IntervalListRequest) (*EventList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MonthlyEventList not implemented")
}
func (UnimplementedEventsServer) mustEmbedUnimplementedEventsServer() {}

// UnsafeEventsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventsServer will
// result in compilation errors.
type UnsafeEventsServer interface {
	mustEmbedUnimplementedEventsServer()
}

func RegisterEventsServer(s grpc.ServiceRegistrar, srv EventsServer) {
	s.RegisterService(&Events_ServiceDesc, srv)
}

func _Events_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Events/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Events/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).UpdateEvent(ctx, req.(*EventActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Events/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).DeleteEvent(ctx, req.(*EventActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_DailyEventList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntervalListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).DailyEventList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Events/DailyEventList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).DailyEventList(ctx, req.(*IntervalListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_WeeklyEventList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntervalListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).WeeklyEventList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Events/WeeklyEventList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).WeeklyEventList(ctx, req.(*IntervalListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_MonthlyEventList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntervalListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).MonthlyEventList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.Events/MonthlyEventList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).MonthlyEventList(ctx, req.(*IntervalListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Events_ServiceDesc is the grpc.ServiceDesc for Events service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Events_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _Events_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _Events_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _Events_DeleteEvent_Handler,
		},
		{
			MethodName: "DailyEventList",
			Handler:    _Events_DailyEventList_Handler,
		},
		{
			MethodName: "WeeklyEventList",
			Handler:    _Events_WeeklyEventList_Handler,
		},
		{
			MethodName: "MonthlyEventList",
			Handler:    _Events_MonthlyEventList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "EventService.proto",
}