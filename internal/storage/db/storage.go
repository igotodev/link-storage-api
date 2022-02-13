package db

import (
	"context"
	"database/sql"
	"link-storage-api/internal/storage/model"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateLink(ctx context.Context, link model.Link) (int, error) {

	row := s.db.QueryRow(`INSERT INTO links (name, category, url, date) VALUES ($1, $2, $3, $4) RETURNING id`,
		link.Name, link.Category, link.URL, link.Date)

	if err := row.Err(); err != nil {
		return 0, err
	}

	var retId int
	if err := row.Scan(&retId); err != nil {
		return 0, err
	}

	return retId, nil
}
func (s *Storage) SelectLink(ctx context.Context, id int) (model.Link, error) {
	return model.Link{}, nil
}
func (s *Storage) SelectAllLinks(ctx context.Context) ([]model.Link, error) {

	query, err := s.db.Query(`SELECT * FROM links`)
	if err != nil {
		return nil, err
	}

	var links []model.Link

	for query.Next() {
		var link model.Link

		err := query.Scan(&link.ID, &link.Name, &link.Category, &link.URL, &link.Date)
		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}
func (s *Storage) UpdateLink(ctx context.Context, id int) error {
	return nil
}
func (s *Storage) DeleteLink(ctx context.Context, id int) error {
	return nil
}
