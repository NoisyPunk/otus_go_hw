package queue

import (
	"github.com/google/uuid"
	"time"
)

type RmqMessage struct {
	EventId     uuid.UUID `json:"eventId"`
	Title       string    `json:"title"`
	DateAndTime time.Time `json:"dateAndTime"`
	UserId      uuid.UUID `json:"userId"`
}
