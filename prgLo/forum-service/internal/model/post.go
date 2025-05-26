package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID
	Title     string
	Content   string
	AuthorID  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostCreate struct {
	Title    string    `json:"title" validate:"required,min=3,max=255"`
	Content  string    `json:"content" validate:"required,min=1"`
	AuthorID uuid.UUID `json:"author_id" validate:"required"`
}

type PostResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  uuid.UUID `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
