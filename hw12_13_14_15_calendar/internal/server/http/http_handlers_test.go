package internalhttp

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/calendar_config"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	appmock "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateEventHandler(t *testing.T) {
	config := &calendarconfig.Config{
		Host: "127.0.0.1",
		Port: "8182",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedApp := appmock.NewMockApplication(mockCtrl)
	server := NewServer(context.Background(), mockedApp, config, nil)
	response := storage.Event{
		ID:           uuid.New(),
		Title:        "test",
		DateAndTime:  time.Now(),
		Duration:     1,
		Description:  "test",
		UserID:       uuid.New(),
		TimeToNotify: 1,
	}

	mockedApp.EXPECT().CreateEvent(gomock.Any(), gomock.Any(), gomock.Any()).Return(response, nil).AnyTimes()

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
			"http://127.0.0.1:8182/create",
			bytes.NewBufferString(`{"title": "test", "dateAndTime": "2023-12-04T12:25:04Z",
    				"duration": 30, "description": "test description","userId": "4a4d4c1f-0c64-41d6-b918-0857987b0bc5",
					"timeToNotify": 500}`),
			http.StatusOK,
		},
		{
			"empty_body",
			http.MethodPost,
			"http://127.0.0.1:8182/create",
			nil,
			http.StatusBadRequest,
		},
		{
			"wrong_method",
			http.MethodGet,
			"http://127.0.0.1:8182/create",
			nil,
			http.StatusMethodNotAllowed,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			r := httptest.NewRequest(c.method, c.target, c.body)
			w := httptest.NewRecorder()
			server.CreateEvent(w, r)
			result := w.Result()
			defer result.Body.Close()
			require.Equal(t, c.responseCode, result.StatusCode)
		})
	}
}

func TestCreateUpdateHandler(t *testing.T) {
	config := &calendarconfig.Config{
		Host: "127.0.0.1",
		Port: "8182",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedApp := appmock.NewMockApplication(mockCtrl)
	server := NewServer(context.Background(), mockedApp, config, nil)

	mockedApp.EXPECT().UpdateEvent(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

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
			"http://127.0.0.1:8182/update",
			bytes.NewBufferString(`{"title": "test", "dateAndTime": "2023-12-04T12:25:04Z",
    				"duration": 30, "description": "test description","userId": "4a4d4c1f-0c64-41d6-b918-0857987b0bc5",
					"timeToNotify": 500}`),
			http.StatusOK,
		},
		{
			"empty_body",
			http.MethodPost,
			"http://127.0.0.1:8182/update",
			nil,
			http.StatusBadRequest,
		},
		{
			"wrong_method",
			http.MethodGet,
			"http://127.0.0.1:8182/update",
			nil,
			http.StatusMethodNotAllowed,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			r := httptest.NewRequest(c.method, c.target, c.body)
			w := httptest.NewRecorder()
			server.UpdateEvent(w, r)
			result := w.Result()
			defer result.Body.Close()
			require.Equal(t, c.responseCode, result.StatusCode)
		})
	}
}
