package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/LizaMeytner/foru/forum-service/internal/model"
	"github.com/google/uuid"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) Create(ctx context.Context, message *model.ChatMessage) error {
	query := `
		INSERT INTO chat_messages (id, content, author_id, created_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(ctx, query,
		message.ID,
		message.Content,
		message.AuthorID,
		time.Now(),
	)

	return err
}

func (r *ChatRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.ChatMessage, error) {
	query := `
		SELECT id, content, author_id, created_at
		FROM chat_messages
		WHERE id = $1
	`

	message := &model.ChatMessage{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&message.ID,
		&message.Content,
		&message.AuthorID,
		&message.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (r *ChatRepository) GetAll(ctx context.Context, offset, limit int) ([]*model.ChatMessage, error) {
	query := `
		SELECT id, content, author_id, created_at
		FROM chat_messages
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*model.ChatMessage
	for rows.Next() {
		message := &model.ChatMessage{}
		err := rows.Scan(
			&message.ID,
			&message.Content,
			&message.AuthorID,
			&message.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *ChatRepository) DeleteOldMessages(ctx context.Context, olderThan time.Time) error {
	query := `DELETE FROM chat_messages WHERE created_at < $1`

	_, err := r.db.ExecContext(ctx, query, olderThan)
	return err
}
