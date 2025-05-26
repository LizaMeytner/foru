package usecase

import (
	"context"

	"github.com/LizaMeytner/foru/forum-service/internal/model"
	"github.com/LizaMeytner/foru/forum-service/internal/repository"
	"github.com/google/uuid"
)

type postUseCase struct {
	postRepo repository.PostRepository
}

func NewPostUseCase(postRepo repository.PostRepository) PostUseCase {
	return &postUseCase{
		postRepo: postRepo,
	}
}

func (u *postUseCase) Create(ctx context.Context, post *model.Post) error {
	return u.postRepo.Create(ctx, post)
}

func (u *postUseCase) GetByID(ctx context.Context, id uuid.UUID) (*model.Post, error) {
	return u.postRepo.GetByID(ctx, id)
}

func (u *postUseCase) GetAll(ctx context.Context, offset, limit int) ([]*model.Post, error) {
	return u.postRepo.GetAll(ctx, offset, limit)
}

func (u *postUseCase) Update(ctx context.Context, post *model.Post) error {
	return u.postRepo.Update(ctx, post)
}

func (u *postUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.postRepo.Delete(ctx, id)
}
