package sqlstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // for DB connection
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrDBConnection = fmt.Errorf("cannot open pgx driver: ")
	ErrCreateEvent  = fmt.Errorf("cannot create event: ")
)

type Storage struct {
	db *sqlx.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, dsn string) (err error) {
	s.db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return errors.Wrap(err, ErrDBConnection.Error())
	}
	return s.db.PingContext(ctx)
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) Create(ctx context.Context, data storage.Event, userID uuid.UUID) (storage.Event, error) {
	l := logger.FromContext(ctx)

	query := `INSERT INTO events (id, user_id, title, date_and_time, duration, description, time_to_notify) 
				VALUES(:id, :user_id, :title, :date_and_time, :duration, :description, :time_to_notify)`

	eventID := uuid.New()
	event := storage.Event{
		ID:           eventID,
		Title:        data.Title,
		DateAndTime:  data.DateAndTime,
		Duration:     data.Duration,
		Description:  data.Description,
		UserID:       userID,
		TimeToNotify: data.TimeToNotify,
	}

	_, err := s.db.NamedQuery(query, event)
	if err != nil {
		l.Error(err.Error(), zap.String("event_id:", eventID.String()))
		return storage.Event{}, errors.Wrap(err, ErrCreateEvent.Error())
	}
	l.Info("event created:", zap.String("event_id", eventID.String()))
	return event, nil
}

func (s *Storage) Update(ctx context.Context, eventID uuid.UUID, event storage.Event) error {
	l := logger.FromContext(ctx)

	query := `UPDATE events SET (user_id, title, date_and_time, duration, description, time_to_notify) =
				(:user_id, :title, :date_and_time, :duration, :description, :time_to_notify) WHERE id=:id`

	_, err := s.db.NamedQuery(query, &event)
	if err != nil {
		l.Error(err.Error(), zap.String("event_id:", eventID.String()))
		return err
	}
	l.Info("event updated:", zap.String("event_id", eventID.String()))
	return nil
}

func (s *Storage) Delete(ctx context.Context, eventID uuid.UUID) error {
	l := logger.FromContext(ctx)

	query := `DELETE FROM events WHERE id=$1`

	_, err := s.db.Exec(query, eventID)
	if err != nil {
		l.Error(err.Error(), zap.String("event_id:", eventID.String()))
		return err
	}
	l.Info("event deleted:", zap.String("event_id", eventID.String()))
	return nil
}

func (s *Storage) DailyList(ctx context.Context, date time.Time, userID uuid.UUID) ([]storage.Event, error) {
	l := logger.FromContext(ctx)

	var events []storage.Event
	query := `SELECT * FROM events where user_id = $1 and date_and_time between $2 and $2 + INTERVAL '1 day'`

	err := s.db.Select(&events, query, userID, date)
	if err != nil {
		return nil, err
	}
	l.Info("dayly list generated:", zap.String("user_id", userID.String()))
	return events, nil
}

func (s *Storage) WeeklyList(ctx context.Context, startWeekDate time.Time, userID uuid.UUID) ([]storage.Event, error) {
	l := logger.FromContext(ctx)

	var events []storage.Event
	query := `SELECT * FROM events where user_id = $1 and date_and_time between $2 and $2 + INTERVAL '1 week'`

	err := s.db.Select(&events, query, userID, startWeekDate)
	if err != nil {
		return nil, err
	}
	l.Info("weekly list generated:", zap.String("user_id", userID.String()))
	return events, nil
}

func (s *Storage) MonthlyList(ctx context.Context, startMonthDate time.Time,
	userID uuid.UUID,
) ([]storage.Event, error) {
	l := logger.FromContext(ctx)
	var events []storage.Event
	query := `SELECT * FROM events where user_id = $1 and date_and_time between $2 and $2 + INTERVAL '1 month'`

	err := s.db.Select(&events, query, userID, startMonthDate)
	if err != nil {
		return nil, err
	}
	l.Info("monthly list generated:", zap.String("user_id", userID.String()))
	return events, nil
}

func (s *Storage) DeleteOldEvents(ctx context.Context, storagePeriod int) error {
	l := logger.FromContext(ctx)
	var events []storage.Event

	period := time.Now().Add(time.Duration(-storagePeriod) * (24 * time.Hour))

	query := `SELECT id FROM events where date_and_time < $1`

	err := s.db.Select(&events, query, period)
	if err != nil {
		return err
	}
	for _, event := range events {
		err = s.Delete(ctx, event.ID)
		if err != nil {
			l.Error("can't delete old event", zap.String("error_message", err.Error()))
		}
	}
	return nil
}

func (s *Storage) NotifyList(ctx context.Context) ([]storage.Event, error) {
	l := logger.FromContext(ctx)
	var events []storage.Event

	query := `SELECT id, title, date_and_time, user_id, time_to_notify
			FROM events where  date_and_time > now() + make_interval(hours := 3) 
              and now() + make_interval(hours := 3) + make_interval(mins := time_to_notify) > date_and_time`

	err := s.db.Select(&events, query)
	if err != nil {
		return nil, err
	}
	l.Info("notify events list generated")
	return events, nil
}
