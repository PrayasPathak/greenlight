package data

import "time"

type Movie struct {
	Title     string    `json:"title"`
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Year      int32     `json:"year,omitempty"`
	Runtime   int32     `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
