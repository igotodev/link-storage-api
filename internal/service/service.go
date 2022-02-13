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

func (s *Service) UpdateLink(ctx context.Context, id int) error {
	return s.storage.UpdateLink(ctx, id)
}

func (s *Service) DeleteLink(ctx context.Context, id int) error {
	return s.storage.DeleteLink(ctx, id)
}
