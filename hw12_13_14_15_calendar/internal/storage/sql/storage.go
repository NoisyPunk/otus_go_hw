package sqlstorage

import (
	"context"
	"fmt"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
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

func (s *Storage) Close(ctx context.Context) error {
	return s.db.Close()
}

func (s *Storage) Create(ctx context.Context, data storage.Event, userID uuid.UUID) (uuid.UUID, error) {
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
		return uuid.UUID{}, errors.Wrap(err, ErrCreateEvent.Error())
	}
	l.Info("event created:", zap.String("event_id", eventID.String()))
	return eventID, nil
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
	var events []storage.Event
	query := `SELECT * FROM events where user_id = $1 and date_and_time between $2 and $2 + INTERVAL '1 day'`

	err := s.db.Select(&events, query, userID, date)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *Storage) WeeklyList(ctx context.Context, startWeekDate time.Time, userID uuid.UUID) ([]storage.Event, error) {
	var events []storage.Event
	query := `SELECT * FROM events where user_id = $1 and date_and_time between $2 and $2 + INTERVAL '1 week'`

	err := s.db.Select(&events, query, userID, startWeekDate)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *Storage) MonthlyList(ctx context.Context, startMonthDate time.Time, userID uuid.UUID) ([]storage.Event, error) {
	var events []storage.Event
	query := `SELECT * FROM events where user_id = $1 and date_and_time between $2 and $2 + INTERVAL '1 month'`

	err := s.db.Select(&events, query, userID, startMonthDate)
	if err != nil {
		return nil, err
	}
	return events, nil
}
