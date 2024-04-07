package memorystorage

import (
	"context"
	"testing"
	"time"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	t.Run("create event", func(t *testing.T) {
		log := logger.New("debug")
		ctx := logger.ContextLogger(context.Background(), log)

		memStorage := New()

		userID := uuid.New()

		event := storage.Event{
			Title:        "Test Title",
			DateAndTime:  time.Date(2024, 11, 5, 23, 24, 23, 0, time.UTC),
			Duration:     15,
			Description:  "Test Decription",
			UserID:       userID,
			TimeToNotify: 10,
		}

		id, err := memStorage.Create(ctx, event, userID)
		require.NotNil(t, id)

		event.ID = id.ID
		require.NoError(t, err)
		require.Equal(t, memStorage.storage[id.ID], event)
	})

	t.Run("create daily list", func(t *testing.T) {
		log := logger.New("debug")
		ctx := logger.ContextLogger(context.Background(), log)

		memStorage := New()

		userID := uuid.New()

		expectedEvents := []storage.Event{
			{
				Title:        "Test Title",
				DateAndTime:  time.Date(2024, 11, 5, 23, 24, 23, 0, time.UTC),
				Duration:     15,
				Description:  "Test Decription",
				UserID:       userID,
				TimeToNotify: 10,
			},
			{
				Title:        "Test Title2",
				DateAndTime:  time.Date(2024, 11, 5, 22, 24, 23, 0, time.UTC),
				Duration:     10,
				Description:  "Test Decription2",
				UserID:       userID,
				TimeToNotify: 14,
			},
		}

		notExpectedEvent := storage.Event{
			Title:        "Test Title3",
			DateAndTime:  time.Date(2024, 11, 8, 22, 24, 23, 0, time.UTC),
			Duration:     10,
			Description:  "Test Decription3",
			UserID:       userID,
			TimeToNotify: 14,
		}

		for _, event := range expectedEvents {
			_, err := memStorage.Create(ctx, event, userID)
			require.NoError(t, err)
		}
		_, err := memStorage.Create(ctx, notExpectedEvent, userID)
		require.NoError(t, err)

		date := time.Date(2024, 11, 5, 22, 24, 23, 0, time.UTC)
		events, err := memStorage.DailyList(ctx, date, userID)
		require.NoError(t, err)

		diffOpt := cmpopts.IgnoreFields(
			storage.Event{}, "ID",
		)

		diff := cmp.Diff(expectedEvents, events, diffOpt)
		require.EqualValues(t, "", diff)
	})
}
