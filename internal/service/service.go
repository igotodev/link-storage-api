package service

import (
	"context"
	"link-storage-api/internal/storage"
	"link-storage-api/internal/storage/model"
)

type Service struct {
	storage storage.StorageImpl
}

func NewService(storage storage.StorageImpl) *Service {
	return &Service{storage: storage}
}

func (s *Service) AddLink(ctx context.Context, link model.Link) (int, error) {
	return s.storage.CreateLink(ctx, link)
}

func (s *Service) Link(ctx context.Context, id int) (model.Link, error) {
	return s.storage.SelectLink(ctx, id)
}

func (s *Service) AllLinks(ctx context.Context) ([]model.Link, error) {
	return s.storage.SelectAllLinks(ctx)
}

func (s *Service) UpdateLink(ctx context.Context, link model.Link) (int, error) {
	return s.storage.UpdateLink(ctx, link)
}

func (s *Service) DeleteLink(ctx context.Context, id int) error {
	return s.storage.DeleteLink(ctx, id)
}

func (s *Service) AddUser(ctx context.Context, user model.User) (int, error) {
	return s.storage.CreateUser(ctx, user)
}

func (s *Service) User(ctx context.Context, username string, passwordHash string) (model.User, error) {
	return s.storage.SelectUser(ctx, username, passwordHash)
}
