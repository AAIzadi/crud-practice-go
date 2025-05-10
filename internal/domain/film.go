package domain

import (
	"crud-practice-go/internal/search"
	"time"
)

type Film struct {
	FilmID          int        `json:"film_id"`
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	ReleaseYear     int        `json:"release_year"`
	LanguageID      int        `json:"language_id"`
	OriginalLangID  *int       `json:"original_language_id"`
	RentalDuration  int        `json:"rental_duration"`
	RentalRate      float64    `json:"rental_rate"`
	Length          int        `json:"length"`
	ReplacementCost float64    `json:"replacement_cost"`
	Rating          string     `json:"rating"`
	LastUpdate      *time.Time `json:"last_update"`
	SpecialFeatures []string   `json:"special_features"`
	FullText        string     `json:"fulltext"`
}

type FilmRepository interface {
	GetAll(param search.PagingAndSorting) ([]Film, error)
	GetById(id int) (film *Film, err error)
	GetFilmsWithLanguage() ([]FilmWithLanguage, error)
}

type FilmWithLanguage struct {
	Title        string
	LanguageName string
}
