package internalhttp

import (
	"time"

	"github.com/google/uuid"
)

type CreateEventRequest struct {
	Title        string        `json:"title"`
	DateAndTime  time.Time     `json:"dateAndTime"`
	Duration     time.Duration `json:"duration"`
	Description  string        `json:"description"`
	UserID       uuid.UUID     `json:"userId"`
	TimeToNotify time.Duration `json:"timeToNotify"`
}

type CreateEventResponse struct {
	Message      string        `json:"message"`
	ID           uuid.UUID     `json:"id"`
	Title        string        `json:"title"`
	DateAndTime  time.Time     `json:"dateAndTime"`
	Duration     time.Duration `json:"duration"`
	Description  string        `json:"description"`
	UserID       uuid.UUID     `json:"userId"`
	TimeToNotify time.Duration `json:"timeToNotify"`
}

type UpdateEventRequest struct {
	EventID uuid.UUID          `json:"eventId"`
	Event   CreateEventRequest `json:"event"`
}

type UpdateEventResponse struct {
	Message string    `json:"message"`
	EventID uuid.UUID `json:"eventId"`
}

type DeleteEventRequest struct {
	EventID uuid.UUID `json:"eventId"`
}

type DeleteEventResponse struct {
	Message string    `json:"message"`
	EventID uuid.UUID `json:"eventId"`
}

type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

type EventListRequest struct {
	DateAndTime time.Time `json:"dateAndTime"`
	UserID      uuid.UUID `json:"userId"`
}

type EventListResponse struct {
	EventList []*CreateEventResponse `json:"eventList"`
}
