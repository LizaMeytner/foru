package usecase

import (
	"context"

	"github.com/LizaMeytner/foru/forum-service/internal/model"
	"github.com/google/uuid"
)

type PostUseCase interface {
	Create(ctx context.Context, post *model.Post) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Post, error)
	GetAll(ctx context.Context, offset, limit int) ([]*model.Post, error)
	Update(ctx context.Context, post *model.Post) error
	Delete(ctx context.Context, id uuid.UUID) error
}
