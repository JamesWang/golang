package data

import (
	"time"
)

type Movie struct {
	ID        int64     `Yaml:"id" json:"id,omitempty"`
	CreatedAt time.Time `Yaml:"createdAt" json:"createdAt,omitempty"`
	Title     string    `Yaml:"title" json:"title,omitempty"`
	Year      int32     `Yaml:"year" json:"year,omitempty"`
	Runtime   int32     `Yaml:"runtime" json:"runtime"`
	Genres    []string  `Yaml:"genres" json:"genres"`
	Version   int32     `Yaml:"version"`
}

func SampleData(id int64) Movie {
	return Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}
}
