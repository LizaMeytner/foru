package usecase

import (
	"context"
	"time"

	"github.com/LizaMeytner/foru/forum-service/internal/model"
	"github.com/LizaMeytner/foru/forum-service/internal/repository"
	"github.com/google/uuid"
)

type chatUseCase struct {
	chatRepo repository.ChatRepository
}

func NewChatUseCase(chatRepo repository.ChatRepository) ChatUseCase {
	return &chatUseCase{
		chatRepo: chatRepo,
	}
}

func (u *chatUseCase) Create(ctx context.Context, message *model.ChatMessage) error {
	return u.chatRepo.Create(ctx, message)
}

func (u *chatUseCase) GetByID(ctx context.Context, id uuid.UUID) (*model.ChatMessage, error) {
	return u.chatRepo.GetByID(ctx, id)
}

func (u *chatUseCase) GetAll(ctx context.Context, offset, limit int) ([]*model.ChatMessage, error) {
	return u.chatRepo.GetAll(ctx, offset, limit)
}

func (u *chatUseCase) DeleteOldMessages(ctx context.Context, olderThan time.Time) error {
	return u.chatRepo.DeleteOldMessages(ctx, olderThan)
}
