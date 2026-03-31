package service

import (
	"strings"
	"article/internal/domain"
	"article/internal/repository"
)

type PostService interface {
	Create(post domain.Post) error
	GetAll() ([]domain.Post, error)
	GetWithPagination(limit, offset int) ([]domain.Post, error)
	GetByID(id int) (domain.Post, error)
	Update(id int, post domain.Post) error
	Delete(id int) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo}
}

func (s *postService) Create(post domain.Post) error {
	if post.Status == "" {
		post.Status = "Draft"
	}
	return s.repo.Create(post)
}

func (s *postService) GetAll() ([]domain.Post, error) {
	return s.repo.GetAll()
}

func (s *postService) GetWithPagination(limit, offset int) ([]domain.Post, error) {
    return s.repo.GetWithPagination(limit, offset)
}

func (s *postService) GetByID(id int) (domain.Post, error) {
    return s.repo.GetByID(id)
}

func (s *postService) Update(id int, post domain.Post) error {
    post.Status = strings.ToLower(post.Status)
    return s.repo.Update(id, post)
}

func (s *postService) Delete(id int) error {
    return s.repo.Delete(id)
}
