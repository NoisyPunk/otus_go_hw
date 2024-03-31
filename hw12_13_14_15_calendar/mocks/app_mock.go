// Code generated by MockGen. DO NOT EDIT.
// Source: app.go

// Package appmock is a generated GoMock package.
package appmock

import (
	context "context"
	reflect "reflect"
	time "time"

	storage "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// CreateEvent mocks base method.
func (m *MockApplication) CreateEvent(ctx context.Context, event storage.Event, userID uuid.UUID) (storage.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", ctx, event, userID)
	ret0, _ := ret[0].(storage.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockApplicationMockRecorder) CreateEvent(ctx, event, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockApplication)(nil).CreateEvent), ctx, event, userID)
}

// DeleteEvent mocks base method.
func (m *MockApplication) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEvent", ctx, eventID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEvent indicates an expected call of DeleteEvent.
func (mr *MockApplicationMockRecorder) DeleteEvent(ctx, eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockApplication)(nil).DeleteEvent), ctx, eventID)
}

// EventsDailyList mocks base method.
func (m *MockApplication) EventsDailyList(ctx context.Context, date time.Time, userID uuid.UUID) ([]storage.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsDailyList", ctx, date, userID)
	ret0, _ := ret[0].([]storage.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EventsDailyList indicates an expected call of EventsDailyList.
func (mr *MockApplicationMockRecorder) EventsDailyList(ctx, date, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsDailyList", reflect.TypeOf((*MockApplication)(nil).EventsDailyList), ctx, date, userID)
}

// EventsMonthlyList mocks base method.
func (m *MockApplication) EventsMonthlyList(ctx context.Context, startMonthDate time.Time, userID uuid.UUID) ([]storage.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsMonthlyList", ctx, startMonthDate, userID)
	ret0, _ := ret[0].([]storage.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EventsMonthlyList indicates an expected call of EventsMonthlyList.
func (mr *MockApplicationMockRecorder) EventsMonthlyList(ctx, startMonthDate, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsMonthlyList", reflect.TypeOf((*MockApplication)(nil).EventsMonthlyList), ctx, startMonthDate, userID)
}

// EventsWeeklyList mocks base method.
func (m *MockApplication) EventsWeeklyList(ctx context.Context, startWeekDate time.Time, userID uuid.UUID) ([]storage.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsWeeklyList", ctx, startWeekDate, userID)
	ret0, _ := ret[0].([]storage.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EventsWeeklyList indicates an expected call of EventsWeeklyList.
func (mr *MockApplicationMockRecorder) EventsWeeklyList(ctx, startWeekDate, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsWeeklyList", reflect.TypeOf((*MockApplication)(nil).EventsWeeklyList), ctx, startWeekDate, userID)
}

// UpdateEvent mocks base method.
func (m *MockApplication) UpdateEvent(ctx context.Context, eventID uuid.UUID, event storage.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEvent", ctx, eventID, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEvent indicates an expected call of UpdateEvent.
func (mr *MockApplicationMockRecorder) UpdateEvent(ctx, eventID, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEvent", reflect.TypeOf((*MockApplication)(nil).UpdateEvent), ctx, eventID, event)
}