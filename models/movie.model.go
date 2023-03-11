package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Name        string    `json:"name"`
	Year        int       `json:"year"`
	Genre       string    `json:"genre"`
	Sinopsis    string    `json:"sinopsis"`
	ReleaseDate time.Time `json:"release_date"`
}
