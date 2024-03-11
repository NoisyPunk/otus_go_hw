package app

import (
	"context"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/google/uuid"
)

type App struct {
	Storage storage.Storage
}

func New(ctx context.Context, config *configs.Config) *App {
	l := logger.FromContext(ctx)
	var store storage.Storage

	switch {
	case config.InmemStore:
		l.Debug("inmem storage is used for server")
		store = memorystorage.New()
	default:
		l.Debug("database storage is used for server")
		store = sqlstorage.New()
	}

	return &App{
		Storage: store,
	}
}

func (a *App) CreateEvent(ctx context.Context, event storage.Event, userID uuid.UUID) (uuid.UUID, error) {
	return a.Storage.Create(ctx, event, userID)
}
