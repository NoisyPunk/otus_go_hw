package internalhttp

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
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
			r, err := http.Post(c.target, "application/json", c.body)
			require.NoError(t, err)
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
			r, err := http.Post(c.target, "application/json", c.body)
			require.NoError(t, err)
			require.Equal(t, c.responseCode, r.StatusCode)
		})
	}
}
