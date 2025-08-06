package service

import (
	"blog/domains/posts/errors"
	"blog/domains/posts/models"
	"blog/domains/posts/repository"

	"github.com/google/uuid"
)

type PostsServiceImpl struct {
	repository *repository.PostsRepository
}

func CreatePostsService(repo *repository.PostsRepository) PostsService {
	return &PostsServiceImpl{repository: repo}
}

func (ps PostsServiceImpl) CreatePost(post *models.Post) (*models.PostResponse, error) {
	if post.Title == "" {
		return nil, errors.ErrorPostWithoutTitle{}
	}
	if post.Content == "" {
		return nil, errors.ErrorPostWithoutContent{}
	}

	response, err := ps.repository.CreatePost(*post)

	if err != nil {
		return nil, err
	}

	return response, err
}

func (ps PostsServiceImpl) UpdatePost(id uuid.UUID, post *models.Post) (*models.PostResponse, error) {

	actual, err := ps.GetPost(id)

	if err != nil {
		return nil, err
	}

	if post.Title == "" {
		post.Title = actual.Title
	}
	if post.Content == "" {
		post.Content = actual.Content
	}

	response, err := ps.repository.UpdatePost(id, *post)

	if err != nil {
		return nil, err
	}

	return response, err
}

func (ps PostsServiceImpl) GetPosts() ([]models.PostResponse, error) {
	response, err := ps.repository.GetPosts()

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (ps PostsServiceImpl) GetPost(id uuid.UUID) (*models.PostResponse, error) {
	response, err := ps.repository.GetPost(id)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (ps PostsServiceImpl) DeletePost(id uuid.UUID) error {
	return ps.repository.DeletePost(id)
}
