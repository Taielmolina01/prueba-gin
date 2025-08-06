package repository

import (
	"blog/domains/posts/errors"
	"blog/domains/posts/models"
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

type PostsRepository struct {
	db *pg.DB
}

func CreatePostsRepository(db *pg.DB) (*PostsRepository, error) {
	schema := `
		CREATE TABLE IF NOT EXISTS posts (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			title VARCHAR(50),
			content VARCHAR(1200),
			publish_date DATE DEFAULT (CURRENT_DATE)
		)
	`

	if _, err := db.Exec(schema); err != nil {
		return nil, fmt.Errorf("Fail creating posts table: %w", err)
	}

	return &PostsRepository{db: db}, nil
}

func (pr *PostsRepository) CreatePost(post models.Post) (*models.PostResponse, error) {
	query := `
		INSERT INTO 
			posts (title, content) 
		VALUES
			(?, ?)
		RETURNING id, title, content, publish_date
	`

	var response models.PostResponse
	_, err := pr.db.QueryOne(&response, query, post.Title, post.Content)
	if err != nil {
		return nil, fmt.Errorf("error inserting post: %w", err)
	}

	return &response, nil
}

func (pr *PostsRepository) UpdatePost(id uuid.UUID, post models.Post) (*models.PostResponse, error) {
	query := `
		UPDATE posts
		SET title = ?, content = ?
		WHERE id = ?
		RETURNING id, title, content, publish_date
	`

	var updatedPost models.PostResponse
	_, err := pr.db.QueryOne(&updatedPost, query, post.Title, post.Content, id)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, errors.ErrorPostDoesntExist{ID: id}
		}
		return nil, fmt.Errorf("Fail updating post: %w", err)
	}

	return &updatedPost, nil
}

func (pr *PostsRepository) GetPosts() ([]models.PostResponse, error) {
	query := `
		SELECT id, title, content, publish_date
		FROM posts
	`

	var posts []models.PostResponse
	_, err := pr.db.Query(&posts, query)
	if err != nil {
		return nil, fmt.Errorf("Fail getting posts: %w", err)
	}

	return posts, nil
}

func (pr *PostsRepository) GetPost(id uuid.UUID) (*models.PostResponse, error) {
	query := `
		SELECT id, title, content, publish_date
		FROM posts
		WHERE id = ?
	`

	var post models.PostResponse
	_, err := pr.db.QueryOne(&post, query, id)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, errors.ErrorPostDoesntExist{ID: id}
		}
		return nil, fmt.Errorf("Fail getting post: %w", err)
	}

	return &post, nil
}

func (pr *PostsRepository) DeletePost(id uuid.UUID) error {
	query := `
		DELETE FROM posts
		WHERE id = ?
	`

	res, err := pr.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Fail deleting post: %w", err)
	}
	if res.RowsAffected() == 0 {
		return errors.ErrorPostDoesntExist{ID: id}
	}

	return nil
}
