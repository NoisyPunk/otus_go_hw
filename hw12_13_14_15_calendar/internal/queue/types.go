package queue

import (
	"time"

	"github.com/google/uuid"
)

type RmqMessage struct {
	EventID     uuid.UUID `json:"eventId"`
	Title       string    `json:"title"`
	DateAndTime time.Time `json:"dateAndTime"`
	UserID      uuid.UUID `json:"userId"`
}
