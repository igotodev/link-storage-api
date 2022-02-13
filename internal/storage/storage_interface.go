package storage

import (
	"context"
	"link-storage-api/internal/storage/model"
)

type StorageImpl interface {
	CreateLink(ctx context.Context, link model.Link) (int, error)
	SelectLink(ctx context.Context, id int) (model.Link, error)
	SelectAllLinks(ctx context.Context) ([]model.Link, error)
	UpdateLink(ctx context.Context, link model.Link) (int, error)
	DeleteLink(ctx context.Context, id int) error
}
