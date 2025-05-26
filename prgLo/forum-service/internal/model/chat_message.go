package model

import (
	"time"

	"github.com/google/uuid"
)

type ChatMessage struct {
	ID        uuid.UUID
	Content   string
	AuthorID  uuid.UUID
	CreatedAt time.Time
}
