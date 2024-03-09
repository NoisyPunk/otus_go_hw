package app

import (
	"context"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"time"
)

type App struct {
	Storage Storage
}

type Storage interface {
	Create(ctx context.Context, data storage.Event, userID uuid.UUID) (uuid.UUID, error)
	Update(ctx context.Context, eventID uuid.UUID, event storage.Event) error
	Delete(ctx context.Context, eventID uuid.UUID) error
	DailyList(ctx context.Context, date time.Time, userID uuid.UUID) ([]storage.Event, error)
	WeeklyList(ctx context.Context, startWeekDate time.Time, userID uuid.UUID) ([]storage.Event, error)
	MonthlyList(ctx context.Context, startMonthDate time.Time, userID uuid.UUID) ([]storage.Event, error)
}

func New(storage Storage) *App {
	return &App{
		Storage: storage,
	}
}

func (a *App) CreateEvent(ctx context.Context, event storage.Event, userID uuid.UUID) (uuid.UUID, error) {
	return a.Storage.Create(ctx, event, userID)
}
