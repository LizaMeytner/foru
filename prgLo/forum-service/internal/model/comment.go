package model

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID
	Content   string
	PostID    uuid.UUID
	AuthorID  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommentCreate struct {
	Content  string    `json:"content" validate:"required,min=1"`
	PostID   uuid.UUID `json:"post_id" validate:"required"`
	AuthorID uuid.UUID `json:"author_id" validate:"required"`
}

type CommentResponse struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	PostID    uuid.UUID `json:"post_id"`
	AuthorID  uuid.UUID `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
