package repository

import (
	"context"
	"time"

	"github.com/LizaMeytner/foru/forum-service/internal/model"
	"github.com/google/uuid"
)

type PostRepository interface {
	Create(ctx context.Context, post *model.Post) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Post, error)
	GetAll(ctx context.Context, offset, limit int) ([]*model.Post, error)
	Update(ctx context.Context, post *model.Post) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type CommentRepository interface {
	Create(ctx context.Context, comment *model.Comment) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Comment, error)
	GetByPostID(ctx context.Context, postID uuid.UUID, offset, limit int) ([]*model.Comment, error)
	Update(ctx context.Context, comment *model.Comment) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ChatRepository interface {
	Create(ctx context.Context, message *model.ChatMessage) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.ChatMessage, error)
	GetAll(ctx context.Context, offset, limit int) ([]*model.ChatMessage, error)
	DeleteOldMessages(ctx context.Context, olderThan time.Time) error
}
