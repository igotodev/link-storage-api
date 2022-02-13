package service

import (
	"context"
	"link-storage-api/internal/storage/model"
)

type ServiceImpl interface {
	AddLink(ctx context.Context, link model.Link) (int, error)
	Link(ctx context.Context, id int) (model.Link, error)
	AllLinks(ctx context.Context) ([]model.Link, error)
	UpdateLink(ctx context.Context, id int) error
	DeleteLink(ctx context.Context, id int) error
}
