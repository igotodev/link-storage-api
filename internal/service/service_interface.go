package service

import (
	"context"
	"link-storage-api/internal/storage/model"
)

type ServiceImpl interface {
	LinkServiceImpl
	UserServiceImpl
}

type LinkServiceImpl interface {
	AddLink(ctx context.Context, link model.Link) (int, error)
	Link(ctx context.Context, id int) (model.Link, error)
	AllLinks(ctx context.Context) ([]model.Link, error)
	UpdateLink(ctx context.Context, link model.Link) (int, error)
	DeleteLink(ctx context.Context, id int) error
}

type UserServiceImpl interface {
	AddUser(ctx context.Context, user model.User) (int, error)
	User(ctx context.Context, username string, passwordHash string) (model.User, error)
}
