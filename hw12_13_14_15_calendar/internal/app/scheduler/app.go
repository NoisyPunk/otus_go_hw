package scheduler

import (
	"context"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/scheduler_config"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/queue"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	sqlstorage "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/pkg/errors"
)

type App struct {
	storage  storage.Storage
	producer *queue.Producer
	frequent int
}

func New(ctx context.Context, config *scheduler_config.Config) (*App, error) {
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
		storage:  store,
		producer: producer,
		frequent: config.Frequency,
	}
	return &app, nil
}

//err = ch.Publish(
//	"",
//	"CalendarQueue",
//	false,
//	false,
//	amqp.Publishing{
//		ContentType: "text/plain",
//		Body:        []byte("hello world"),
//	},
//)
//if err != nil {
//	panic(err)
//}
