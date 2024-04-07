package sender

import (
	"context"
	"encoding/json"

	senderconfig "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/sender_config"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/queue"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type App struct {
	consumer *queue.Consumer
}

func New(ctx context.Context, config *senderconfig.Config) (*App, error) {
	consumer, err := queue.NewConsumer(ctx, config)
	if err != nil {
		return nil, errors.Wrap(err, "error creating consumer")
	}

	app := App{
		consumer: consumer,
	}
	return &app, nil
}

func (a *App) Consume(ctx context.Context) {
	l := logger.FromContext(ctx)
	qMsgs, err := a.consumer.RmqChannel.Consume(
		"CalendarQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		l.Error("can't consume channel", zap.String("error_message", err.Error()))
	}

	var msg queue.RmqMessage

	go func() {
		for qMsg := range qMsgs {
			err = json.Unmarshal(qMsg.Body, &msg)
			if err != nil {
				l.Error("can't unmarshal message", zap.String("error_message", err.Error()))
			}
			l.Info("received message:", zap.String("event_id", msg.EventID.String()), zap.String("event_title", msg.Title),
				zap.String("date_and_time", msg.DateAndTime.String()), zap.String("user_id", msg.UserID.String()))
		}
	}()

	<-ctx.Done()
}
