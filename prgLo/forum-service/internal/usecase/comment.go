package usecase

import (
	"context"

	"github.com/LizaMeytner/foru/forum-service/internal/model"
	"github.com/google/uuid"
)

type CommentUseCase interface {
	Create(ctx context.Context, comment *model.Comment) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Comment, error)
	GetByPostID(ctx context.Context, postID uuid.UUID, offset, limit int) ([]*model.Comment, error)
	Update(ctx context.Context, comment *model.Comment) error
	Delete(ctx context.Context, id uuid.UUID) error
}
