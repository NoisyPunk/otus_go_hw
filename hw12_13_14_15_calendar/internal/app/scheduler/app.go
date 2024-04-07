package scheduler

import (
	"context"
	"encoding/json"
	"time"

	schedulerConfig "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/scheduler_config"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/queue"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	sqlstorage "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type App struct {
	storage               storage.Storage
	producer              *queue.Producer
	removeScannerFrequent int
	notifyScannerFrequent int
	storePeriod           int
}

func New(ctx context.Context, config *schedulerConfig.Config) (*App, error) {
	producer, err := queue.NewProducer(ctx, config)
	if err != nil {
		return nil, errors.Wrap(err, "error creating producer")
	}
	store := sqlstorage.New()
	err = store.Connect(ctx, config.Dsn)
	if err != nil {
		return nil, errors.Wrap(err, "error with connecting to storage by producer")
	}
	app := App{
		storage:               store,
		producer:              producer,
		removeScannerFrequent: config.RemoveScannerFrequency,
		notifyScannerFrequent: config.NotifyScannerFrequency,
		storePeriod:           config.StoragePeriod,
	}
	return &app, nil
}

func (a *App) OldEventRemover(ctx context.Context) {
	l := logger.FromContext(ctx)

	ticker := time.NewTicker(time.Duration(a.removeScannerFrequent) * time.Minute)

	for {
		select {
		case <-ticker.C:
			events, err := a.storage.OldEventsList(ctx, a.storePeriod)
			if err != nil {
				l.Error("can't get list of old events", zap.String("error_message", err.Error()))
			}
			for _, event := range events {
				err = a.storage.Delete(ctx, event.ID)
				if err != nil {
					l.Error("can't delete old event", zap.String("error_message", err.Error()))
				}
			}
			l.Info("old events deleted successfully")

		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func (a *App) Notifier(ctx context.Context) {
	l := logger.FromContext(ctx)

	ticker := time.NewTicker(time.Duration(a.notifyScannerFrequent) * time.Minute)

	for {
		select {
		case <-ticker.C:
			events, err := a.storage.NotifyList(ctx)
			if err != nil {
				l.Error("can't get events list for notify", zap.String("error_message", err.Error()))
			}
			for _, event := range events {
				message := queue.RmqMessage{
					EventID:     event.ID,
					Title:       event.Title,
					DateAndTime: event.DateAndTime,
					UserID:      event.UserID,
				}

				j, err := json.Marshal(message)
				if err != nil {
					l.Error("can't marshal event for queue", zap.String("error_message", err.Error()))
				}

				err = a.producer.RmqChannel.Publish(
					"",
					"CalendarQueue",
					false,
					false,
					amqp.Publishing{
						ContentType: "json",
						Body:        j,
					},
				)
				if err != nil {
					l.Error("can't publish event to queue",
						zap.String("event_id", event.ID.String()),
						zap.String("error_message", err.Error()),
					)
				}
			}
			l.Info("events published successfully")

		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}
