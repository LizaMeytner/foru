package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/LizaMeytner/foru/forum-service/internal/model"
	"github.com/google/uuid"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(ctx context.Context, comment *model.Comment) error {
	query := `
		INSERT INTO comments (id, content, post_id, author_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		comment.ID,
		comment.Content,
		comment.PostID,
		comment.AuthorID,
		time.Now(),
		time.Now(),
	)

	return err
}

func (r *CommentRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Comment, error) {
	query := `
		SELECT id, content, post_id, author_id, created_at, updated_at
		FROM comments
		WHERE id = $1
	`

	comment := &model.Comment{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&comment.ID,
		&comment.Content,
		&comment.PostID,
		&comment.AuthorID,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentRepository) GetByPostID(ctx context.Context, postID uuid.UUID, offset, limit int) ([]*model.Comment, error) {
	query := `
		SELECT id, content, post_id, author_id, created_at, updated_at
		FROM comments
		WHERE post_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, postID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*model.Comment
	for rows.Next() {
		comment := &model.Comment{}
		err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.PostID,
			&comment.AuthorID,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentRepository) Update(ctx context.Context, comment *model.Comment) error {
	query := `
		UPDATE comments
		SET content = $1, updated_at = $2
		WHERE id = $3
	`

	_, err := r.db.ExecContext(ctx, query,
		comment.Content,
		time.Now(),
		comment.ID,
	)

	return err
}

func (r *CommentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM comments WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
