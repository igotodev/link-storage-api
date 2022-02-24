package storage

import (
	"context"
	"link-storage-api/internal/storage/model"
)

type StorageImpl interface {
	LinkImpl
	UserImpl
}

type LinkImpl interface {
	CreateLink(ctx context.Context, link model.Link) (int, error)
	SelectLink(ctx context.Context, id int) (model.Link, error)
	SelectAllLinks(ctx context.Context) ([]model.Link, error)
	UpdateLink(ctx context.Context, link model.Link) (int, error)
	DeleteLink(ctx context.Context, id int) error
}

type UserImpl interface {
	CreateUser(ctx context.Context, user model.User) (int, error)
	SelectUser(ctx context.Context, username string, passwordHash string) (model.User, error)
}
