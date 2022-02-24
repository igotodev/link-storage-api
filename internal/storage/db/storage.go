package db

import (
	"context"
	"database/sql"
	"link-storage-api/internal/storage/model"
	"time"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateLink(ctx context.Context, link model.Link) (int, error) {

	row := s.db.QueryRowContext(ctx, `INSERT INTO links (name, category, url, date) VALUES ($1, $2, $3, $4) RETURNING id`,
		link.Name, link.Category, link.URL, time.Now())

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

	row := s.db.QueryRowContext(ctx, `SELECT * FROM links WHERE id=$1`, id)

	if err := row.Err(); err != nil {
		return model.Link{}, err
	}

	var link model.Link

	if err := row.Scan(&link.ID, &link.Name, &link.Category, &link.URL, &link.Date); err != nil {
		return model.Link{}, err
	}

	return link, nil
}
func (s *Storage) SelectAllLinks(ctx context.Context) ([]model.Link, error) {

	query, err := s.db.QueryContext(ctx, `SELECT * FROM links`)
	if err != nil {
		return nil, err
	}

	var links []model.Link

	defer query.Close()

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
func (s *Storage) UpdateLink(ctx context.Context, link model.Link) (int, error) {

	row := s.db.QueryRowContext(ctx, `UPDATE links SET name=$2, category=$3, url=$4, date=$5 WHERE id=$1 RETURNING id`,
		link.ID, link.Name, link.Category, link.URL, time.Now())

	if err := row.Err(); err != nil {
		return 0, err
	}

	var id int

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (s *Storage) DeleteLink(ctx context.Context, id int) error {

	_, err := s.db.ExecContext(ctx, `DELETE FROM links WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateUser(ctx context.Context, user model.User) (int, error) {
	row := s.db.QueryRowContext(ctx, `INSERT INTO users (username, password, active) VALUES ($1, $2, $3) returning id`,
		user.Username, user.PasswordHash, user.Active)

	if err := row.Err(); err != nil {
		return 0, err
	}

	var id int

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) SelectUser(ctx context.Context, username string, passwordHash string) (model.User, error) {
	row := s.db.QueryRowContext(ctx, `SELECT * FROM users WHERE username=$1 AND password=$2`,
		username, passwordHash)

	if err := row.Err(); err != nil {
		return model.User{}, err
	}

	var user model.User

	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Active)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
