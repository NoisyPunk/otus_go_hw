package internalhttp

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	sqlstorage "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateEventHandler(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		target       string
		body         io.Reader
		responseCode int
	}{
		{
			"ok",
			http.MethodPost,
			"http://calendar.service:8182/create",
			bytes.NewBufferString(`{"title": "test", "dateAndTime": "2023-12-04T12:25:04Z",
    				"duration": 30, "description": "test description","userId": "4a4d4c1f-0c64-41d6-b918-0857987b0bc5",
					"timeToNotify": 500}`),
			http.StatusOK,
		},
		{
			"empty_body",
			http.MethodPost,
			"http://calendar.service:8182/create",
			nil,
			http.StatusBadRequest,
		},
		{
			"wrong_method",
			http.MethodGet,
			"http://calendar.service:8182/create",
			nil,
			http.StatusBadRequest,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			r, err := http.Post(c.target, "application/json", c.body) //nolint:noctx
			require.NoError(t, err)
			defer r.Body.Close()
			require.Equal(t, c.responseCode, r.StatusCode)
		})
	}
}

func TestCreateUpdateHandler(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		target       string
		body         io.Reader
		responseCode int
	}{
		{
			"ok",
			http.MethodPost,
			"http://calendar.service:8182/update",
			bytes.NewBufferString(`{"title": "test", "dateAndTime": "2023-12-04T12:25:04Z",
    				"duration": 30, "description": "test description","userId": "4a4d4c1f-0c64-41d6-b918-0857987b0bc5",
					"timeToNotify": 500}`),
			http.StatusOK,
		},
		{
			"empty_body",
			http.MethodPost,
			"http://calendar.service:8182/update",
			nil,
			http.StatusBadRequest,
		},
		{
			"wrong_method",
			http.MethodGet,
			"http://calendar.service:8182/update",
			nil,
			http.StatusBadRequest,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			r, err := http.Post(c.target, "application/json", c.body) //nolint:noctx
			require.NoError(t, err)
			defer r.Body.Close()
			require.Equal(t, c.responseCode, r.StatusCode)
		})
	}
}

func TestEvent(t *testing.T) {
	log := logger.New("debug")
	ctx := logger.ContextLogger(context.Background(), log)

	DBstorage := sqlstorage.New()

	dsn := "host=testdb.local port=5433 user=postgres password=postgres dbname=calendar sslmode=disable"

	err := DBstorage.Connect(ctx, dsn)
	require.NoError(t, err)

	t.Run("create event", func(t *testing.T) {
		userID := uuid.New()

		expectedEvent := storage.Event{
			Title:        "Test Title",
			DateAndTime:  time.Date(2024, 11, 5, 23, 24, 23, 0, time.UTC),
			Duration:     15,
			Description:  "Test Description",
			UserID:       userID,
			TimeToNotify: 10,
		}

		id, err := DBstorage.Create(ctx, expectedEvent, userID)
		require.NotNil(t, id)

		expectedEvent.ID = id.ID
		require.NoError(t, err)

		var events []*storage.Event
		query := `SELECT * FROM events where user_id = $1`
		err = DBstorage.DB.Select(&events, query, expectedEvent.UserID)
		require.NoError(t, err)
		require.Equal(t, &expectedEvent, events[0])
	})
	t.Run("update event", func(t *testing.T) {
		userID := uuid.New()

		expectedEvent := storage.Event{
			Title:        "Test Title",
			DateAndTime:  time.Date(2024, 11, 5, 23, 24, 23, 0, time.UTC),
			Duration:     15,
			Description:  "Test Description",
			UserID:       userID,
			TimeToNotify: 10,
		}

		id, err := DBstorage.Create(ctx, expectedEvent, userID)
		require.NoError(t, err)
		require.NotNil(t, id)

		event := storage.Event{
			ID:           id.ID,
			Title:        "Test Title Updated",
			DateAndTime:  time.Date(2024, 11, 5, 23, 24, 23, 0, time.UTC),
			Duration:     15,
			Description:  "Test Description Updated",
			UserID:       userID,
			TimeToNotify: 10,
		}

		err = DBstorage.Update(ctx, id.ID, event)
		require.NoError(t, err)

		var events []*storage.Event
		query := `SELECT * FROM events where id = $1`
		err = DBstorage.DB.Select(&events, query, id.ID)
		require.NoError(t, err)
		require.Equal(t, &event, events[0])
	})
	t.Run("delete event", func(t *testing.T) {
		userID := uuid.New()

		expectedEvent := storage.Event{
			Title:        "Test Title",
			DateAndTime:  time.Date(2024, 11, 5, 23, 24, 23, 0, time.UTC),
			Duration:     15,
			Description:  "Test Description",
			UserID:       userID,
			TimeToNotify: 10,
		}

		id, err := DBstorage.Create(ctx, expectedEvent, userID)
		require.NoError(t, err)
		require.NotNil(t, id)

		err = DBstorage.Delete(ctx, id.ID)
		require.NoError(t, err)

		query := `SELECT * FROM events where id = $1`
		row, err := DBstorage.DB.Exec(query, expectedEvent.UserID)
		require.NoError(t, err)
		c, err := row.RowsAffected()
		require.NoError(t, err)
		require.Zero(t, c)
	})
}

func TestEventLists(t *testing.T) {
	log := logger.New("debug")
	ctx := logger.ContextLogger(context.Background(), log)

	DBstorage := sqlstorage.New()

	dsn := "host=testdb.local port=5432 user=postgres password=postgres dbname=calendar sslmode=disable"

	err := DBstorage.Connect(ctx, dsn)
	require.NoError(t, err)

	t.Run("daily list", func(t *testing.T) {
		userID := uuid.New()

		events := []storage.Event{
			{
				Title:        "Test Title",
				DateAndTime:  time.Date(2024, 11, 6, 23, 25, 23, 0, time.UTC),
				Duration:     15,
				Description:  "Test Description",
				UserID:       userID,
				TimeToNotify: 10,
			},
			{
				Title:        "Test Title",
				DateAndTime:  time.Date(2024, 11, 5, 23, 24, 23, 0, time.UTC),
				Duration:     15,
				Description:  "Test Description",
				UserID:       userID,
				TimeToNotify: 10,
			},
			{
				Title:        "Test Title",
				DateAndTime:  time.Date(2024, 11, 5, 23, 24, 23, 0, time.UTC),
				Duration:     15,
				Description:  "Test Description",
				UserID:       userID,
				TimeToNotify: 10,
			},
			{
				Title:        "Test Title",
				DateAndTime:  time.Date(2024, 11, 4, 23, 24, 23, 0, time.UTC),
				Duration:     15,
				Description:  "Test Description",
				UserID:       userID,
				TimeToNotify: 10,
			},
		}

		for _, event := range events {
			id, err := DBstorage.Create(ctx, event, userID)
			require.NotNil(t, id)
			require.NoError(t, err)
		}
		date := time.Date(2024, 11, 5, 23, 24, 23, 0, time.UTC)
		result, err := DBstorage.DailyList(ctx, date, userID)
		require.NoError(t, err)
		require.Len(t, result, 2)
	})
	t.Run("weekly list", func(t *testing.T) {
		userID := uuid.New()

		events := []storage.Event{
			{
				Title:        "Test Title",
				DateAndTime:  time.Date(2024, 11, 6, 23, 25, 23, 0, time.UTC),
				Duration:     15,
				Description:  "Test Description",
				UserID:       userID,
				TimeToNotify: 10,
			},
			{
				Title:        "Test Title",
				DateAndTime:  time.Date(2024, 11, 12, 23, 24, 23, 0, time.UTC),
				Duration:     15,
				Description:  "Test Description",
				UserID:       userID,
				TimeToNotify: 10,
			},
			{
				Title:        "Test Title",
				DateAndTime:  time.Date(2024, 11, 14, 23, 24, 23, 0, time.UTC),
				Duration:     15,
				Description:  "Test Description",
				UserID:       userID,
				TimeToNotify: 10,
			},
			{
				Title:        "Test Title",
				DateAndTime:  time.Date(2024, 11, 15, 23, 24, 23, 0, time.UTC),
				Duration:     15,
				Description:  "Test Description",
				UserID:       userID,
				TimeToNotify: 10,
			},
		}

		for _, event := range events {
			id, err := DBstorage.Create(ctx, event, userID)
			require.NotNil(t, id)
			require.NoError(t, err)
		}
		date := time.Date(2024, 11, 7, 23, 24, 23, 0, time.UTC)
		result, err := DBstorage.WeeklyList(ctx, date, userID)
		require.NoError(t, err)
		require.Len(t, result, 2)
	})
}
