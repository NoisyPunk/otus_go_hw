package scheduler

import (
	"context"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/scheduler_config"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/queue"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	storage  storage.Storage
	producer queue.Producer
	frequent int
}

func New(ctx context.Context, config *scheduler_config.Config) (*App, error) {
	producer, err := queue.NewProducer(ctx, config)
	_, _ = producer, err
	return nil, nil
}
