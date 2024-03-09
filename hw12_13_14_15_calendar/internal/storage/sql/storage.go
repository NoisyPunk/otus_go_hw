package sqlstorage

import (
	"context"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"time"
)

type Storage struct { // TODO
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context) error {
	// TODO
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	// TODO
	return nil
}

func (s *Storage) Create(ctx context.Context, data storage.Event, userID uuid.UUID) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (s *Storage) Update(ctx context.Context, eventID uuid.UUID, event storage.Event) error {
	return nil
}

func (s *Storage) Delete(ctx context.Context, eventID uuid.UUID) error {
	return nil
}

func (s *Storage) DailyList(ctx context.Context, date time.Time, userID uuid.UUID) ([]storage.Event, error) {
	return nil, nil
}

func (s *Storage) WeeklyList(ctx context.Context, startWeekDate time.Time, userID uuid.UUID) ([]storage.Event, error) {
	return nil, nil
}

func (s *Storage) MonthlyList(ctx context.Context, startMonthDate time.Time, userID uuid.UUID) ([]storage.Event, error) {
	return nil, nil
}
