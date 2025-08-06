package service

import (
	"blog/domains/posts/models"

	"github.com/google/uuid"
)

type PostsService interface {
	CreatePost(*models.Post) (*models.PostResponse, error)
	UpdatePost(uuid.UUID, *models.Post) (*models.PostResponse, error)
	GetPosts() ([]models.PostResponse, error)
	GetPost(uuid.UUID) (*models.PostResponse, error)
	DeletePost(uuid.UUID) error
}
