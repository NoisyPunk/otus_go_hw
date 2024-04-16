package storage

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID           uuid.UUID `db:"id"`
	Title        string    `db:"title"`
	DateAndTime  time.Time `db:"date_and_time"`
	Duration     int       `db:"duration"`
	Description  string    `db:"description"`
	UserID       uuid.UUID `db:"user_id"`
	TimeToNotify int       `db:"time_to_notify"`
}

type Storage interface {
	Connect(ctx context.Context, dsn string) (err error)
	Close() error
	Create(ctx context.Context, data Event, userID uuid.UUID) (Event, error)
	Update(ctx context.Context, eventID uuid.UUID, event Event) error
	Delete(ctx context.Context, eventID uuid.UUID) error
	DailyList(ctx context.Context, date time.Time, userID uuid.UUID) ([]Event, error)
	WeeklyList(ctx context.Context, startWeekDate time.Time, userID uuid.UUID) ([]Event, error)
	MonthlyList(ctx context.Context, startMonthDate time.Time, userID uuid.UUID) ([]Event, error)
	DeleteOldEvents(limit int) error
	NotifyList(ctx context.Context) ([]Event, error)
}
