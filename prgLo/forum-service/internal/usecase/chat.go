package usecase

import (
	"context"
	"time"

	"github.com/LizaMeytner/foru/forum-service/internal/model"
	"github.com/google/uuid"
)

type ChatUseCase interface {
	Create(ctx context.Context, message *model.ChatMessage) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.ChatMessage, error)
	GetAll(ctx context.Context, offset, limit int) ([]*model.ChatMessage, error)
	DeleteOldMessages(ctx context.Context, olderThan time.Time) error
}
