package model

import "time"

type Link struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	URL      string    `json:"url"`
	Date     time.Time `json:"date"`
}
