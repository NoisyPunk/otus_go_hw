package internalhttp

import (
	"github.com/google/uuid"
	"time"
)

type CreateEventRequest struct {
	Title        string        `json:"title"`
	DateAndTime  time.Time     `json:"date_and_time"`
	Duration     time.Duration `json:"duration"`
	Description  string        `json:"description"`
	UserID       uuid.UUID     `json:"user_id"`
	TimeToNotify time.Duration `json:"time_to_notify"`
}

type CreateEventResponse struct {
	Message      string        `json:"message"`
	ID           uuid.UUID     `json:"id"`
	Title        string        `json:"title"`
	DateAndTime  time.Time     `json:"date_and_time"`
	Duration     time.Duration `json:"duration"`
	Description  string        `json:"description"`
	UserID       uuid.UUID     `json:"user_id"`
	TimeToNotify time.Duration `json:"time_to_notify"`
	Error        struct {
		Message string `json:"message"`
	} `json:"error"`
}

type UpdateEventRequest struct {
	EventID uuid.UUID          `json:"event_id"`
	Event   CreateEventRequest `json:"event"`
}

type UpdateEventResponse struct {
	Message string    `json:"message"`
	EventID uuid.UUID `json:"event_id"`
	Error   struct {
		Message string `json:"message"`
	} `json:"error"`
}

type DeleteEventRequest struct {
	EventID uuid.UUID `json:"event_id"`
}

type DeleteEventResponse struct {
	Message string    `json:"message"`
	EventID uuid.UUID `json:"event_id"`
	Error   struct {
		Message string `json:"message"`
	} `json:"error"`
}
