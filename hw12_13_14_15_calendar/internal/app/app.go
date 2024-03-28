package app

import (
	"context"
	"sync"
	"time"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/google/uuid"
)

type App struct {
	Storage storage.Storage
	mu      sync.RWMutex
}

type Application interface {
	CreateEvent(ctx context.Context, event storage.Event, userID uuid.UUID) (storage.Event, error)
	UpdateEvent(ctx context.Context, eventID uuid.UUID, event storage.Event) error
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
	EventsDailyList(ctx context.Context, date time.Time, userID uuid.UUID) ([]storage.Event, error)
	EventsWeeklyList(ctx context.Context, startWeekDate time.Time, userID uuid.UUID) ([]storage.Event, error)
	EventsMonthlyList(ctx context.Context, startMonthDate time.Time, userID uuid.UUID) ([]storage.Event, error)
}

func New(ctx context.Context, config *configs.Config) (*App, error) {
	l := logger.FromContext(ctx)
	var store storage.Storage

	switch {
	case config.InmemStore:
		l.Debug("inmem storage is used for server")
		store = memorystorage.New()
	default:
		l.Debug("database storage is used for server")
		store = sqlstorage.New()
		err := store.Connect(ctx, config.Dsn)
		if err != nil {
			return nil, err
		}
	}

	return &App{
		Storage: store,
		mu:      sync.RWMutex{},
	}, nil
}

func (a *App) CreateEvent(ctx context.Context, event storage.Event, userID uuid.UUID) (storage.Event, error) {
	return a.Storage.Create(ctx, event, userID)
}

func (a *App) UpdateEvent(ctx context.Context, eventID uuid.UUID, event storage.Event) error {
	return a.Storage.Update(ctx, eventID, event)
}

func (a *App) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	return a.Storage.Delete(ctx, eventID)
}

func (a *App) EventsDailyList(ctx context.Context, date time.Time, userID uuid.UUID) ([]storage.Event, error) {
	return a.Storage.DailyList(ctx, date, userID)
}

func (a *App) EventsWeeklyList(ctx context.Context,
	startWeekDate time.Time, userID uuid.UUID,
) ([]storage.Event, error) {
	return a.Storage.WeeklyList(ctx, startWeekDate, userID)
}

func (a *App) EventsMonthlyList(ctx context.Context,
	startMonthDate time.Time, userID uuid.UUID,
) ([]storage.Event, error) {
	return a.Storage.MonthlyList(ctx, startMonthDate, userID)
}
