package memorystorage

import (
	"context"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"go.uber.org/zap"
	"sync"
)

type Storage struct {
	// TODO
	mu sync.RWMutex //nolint:unused
}

func New(ctx context.Context) *Storage {
	l := logger.FromContext(ctx)

	l.Info("You are in storage Info")
	l.Error("You are in storage Err")
	l.Debug("You are in storage debug")
	l.Warn("You are in storage warn")
	var s bool
	s = false

	if s == false {
		test(ctx)
	}
	return &Storage{}
}

func test(ctx context.Context) {
	l := logger.FromContext(ctx).With(zap.String("Storage", "Storage2"))
	l.Info("Here")
}

// TODO
