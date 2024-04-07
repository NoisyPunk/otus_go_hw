package memorystorage

import (
	"context"
	"sync"
	"time"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Storage struct {
	mu      sync.RWMutex
	storage map[uuid.UUID]storage.Event
}

func New() *Storage {
	return &Storage{
		storage: make(map[uuid.UUID]storage.Event),
	}
}

func (s *Storage) Connect(_ context.Context, _ string) (err error) {
	return errors.New("not implemented for memory storage")
}

func (s *Storage) Close() error {
	return errors.New("not implemented for memory storage")
}

func (s *Storage) Create(ctx context.Context, data storage.Event, userID uuid.UUID) (storage.Event, error) {
	l := logger.FromContext(ctx).With(zap.String("user_id", data.UserID.String()))

	eventID := uuid.New()
	event := storage.Event{
		ID:           eventID,
		Title:        data.Title,
		DateAndTime:  data.DateAndTime,
		Duration:     data.Duration,
		Description:  data.Description,
		UserID:       userID,
		TimeToNotify: data.TimeToNotify,
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.storage[eventID] = event
	l.Info("event created ", zap.String("event_id:", eventID.String()))
	return event, nil
}

func (s *Storage) Update(ctx context.Context, eventID uuid.UUID, event storage.Event) error {
	l := logger.FromContext(ctx)

	s.mu.Lock()
	defer s.mu.Unlock()
	s.storage[eventID] = event
	l.Info("event updated ", zap.String("event_id:", eventID.String()))
	return nil
}

func (s *Storage) Delete(ctx context.Context, eventID uuid.UUID) error {
	l := logger.FromContext(ctx)

	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.storage, eventID)
	l.Info("event deleted", zap.String("event_id:", eventID.String()))
	return nil
}

func (s *Storage) DailyList(ctx context.Context, date time.Time, userID uuid.UUID) ([]storage.Event, error) {
	l := logger.FromContext(ctx)

	events := make([]storage.Event, 0)

	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range s.storage {
		if v.UserID == userID && v.DateAndTime.YearDay() == date.YearDay() {
			events = append(events, v)
		}
	}
	l.Info("daily list formed", zap.String("user_id:", userID.String()))
	return events, nil
}

func (s *Storage) WeeklyList(ctx context.Context, startWeekDate time.Time, userID uuid.UUID) ([]storage.Event, error) {
	l := logger.FromContext(ctx)

	events := make([]storage.Event, 0)

	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range s.storage {
		if v.UserID == userID && startWeekDate.YearDay() <= v.DateAndTime.YearDay() &&
			v.DateAndTime.YearDay() <= startWeekDate.YearDay()+7 {
			events = append(events, v)
		}
	}
	l.Info("weekly list formed", zap.String("user_id:", userID.String()))
	return events, nil
}

func (s *Storage) MonthlyList(ctx context.Context, startMonthDate time.Time,
	userID uuid.UUID,
) ([]storage.Event, error) {
	l := logger.FromContext(ctx)

	events := make([]storage.Event, 0)

	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range s.storage {
		if v.UserID == userID && startMonthDate.YearDay() <= v.DateAndTime.YearDay() &&
			v.DateAndTime.YearDay() <= startMonthDate.YearDay()+30 {
			events = append(events, v)
		}
	}
	l.Info("monthly list formed", zap.String("user_id:", userID.String()))
	return events, nil
}

func (s *Storage) OldEventsList(_ context.Context, _ int) ([]storage.Event, error) {
	return nil, nil
}

func (s *Storage) NotifyList(_ context.Context) ([]storage.Event, error) {
	return nil, nil
}
