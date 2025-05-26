package usecase

import (
	"context"

	"github.com/LizaMeytner/foru/forum-service/internal/model"
	"github.com/LizaMeytner/foru/forum-service/internal/repository"
	"github.com/google/uuid"
)

type commentUseCase struct {
	commentRepo repository.CommentRepository
}

func NewCommentUseCase(commentRepo repository.CommentRepository) CommentUseCase {
	return &commentUseCase{
		commentRepo: commentRepo,
	}
}

func (u *commentUseCase) Create(ctx context.Context, comment *model.Comment) error {
	return u.commentRepo.Create(ctx, comment)
}

func (u *commentUseCase) GetByID(ctx context.Context, id uuid.UUID) (*model.Comment, error) {
	return u.commentRepo.GetByID(ctx, id)
}

func (u *commentUseCase) GetByPostID(ctx context.Context, postID uuid.UUID, offset, limit int) ([]*model.Comment, error) {
	return u.commentRepo.GetByPostID(ctx, postID, offset, limit)
}

func (u *commentUseCase) Update(ctx context.Context, comment *model.Comment) error {
	return u.commentRepo.Update(ctx, comment)
}

func (u *commentUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.commentRepo.Delete(ctx, id)
}
