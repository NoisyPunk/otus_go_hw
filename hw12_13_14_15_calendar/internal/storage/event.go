package storage

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID           uuid.UUID
	Title        string
	DateAndTime  time.Time
	Duration     time.Duration
	Description  string
	UserID       uuid.UUID
	TimeToNotify time.Duration
}
