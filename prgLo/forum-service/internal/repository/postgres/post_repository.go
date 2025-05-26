package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/LizaMeytner/foru/forum-service/internal/model"
	"github.com/google/uuid"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(ctx context.Context, post *model.Post) error {
	query := `
		INSERT INTO posts (id, title, content, author_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		post.ID,
		post.Title,
		post.Content,
		post.AuthorID,
		time.Now(),
		time.Now(),
	)

	return err
}

func (r *PostRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Post, error) {
	query := `
		SELECT id, title, content, author_id, created_at, updated_at
		FROM posts
		WHERE id = $1
	`

	post := &model.Post{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.AuthorID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepository) GetAll(ctx context.Context, offset, limit int) ([]*model.Post, error) {
	query := `
		SELECT id, title, content, author_id, created_at, updated_at
		FROM posts
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*model.Post
	for rows.Next() {
		post := &model.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) Update(ctx context.Context, post *model.Post) error {
	query := `
		UPDATE posts
		SET title = $1, content = $2, updated_at = $3
		WHERE id = $4
	`

	_, err := r.db.ExecContext(ctx, query,
		post.Title,
		post.Content,
		time.Now(),
		post.ID,
	)

	return err
}

func (r *PostRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM posts WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
