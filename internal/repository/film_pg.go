package repository

import (
	"context"
	"crud-practice-go/internal/domain"
	"crud-practice-go/internal/search"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
)

type filmPgRepository struct {
	db *pgxpool.Pool
}

func NewFilmRepository(db *pgxpool.Pool) domain.FilmRepository {
	return &filmPgRepository{db: db}
}

func (r *filmPgRepository) GetAll(param search.PagingAndSorting) ([]domain.Film, error) {
	ctx := context.Background()
	rows, err := r.db.Query(ctx, `SELECT film_id, title, description, release_year, language_id,
        original_language_id, rental_duration, rental_rate, length, replacement_cost,
        rating, last_update, special_features, fulltext FROM film LIMIT $1 OFFSET $2`, param.Limit, param.Offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var films []domain.Film
	for rows.Next() {
		var film domain.Film
		var specialFeatures []byte

		err := rows.Scan(
			&film.FilmID, &film.Title, &film.Description, &film.ReleaseYear,
			&film.LanguageID, &film.OriginalLangID, &film.RentalDuration,
			&film.RentalRate, &film.Length, &film.ReplacementCost,
			&film.Rating, &film.LastUpdate, &specialFeatures, &film.FullText,
		)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(specialFeatures, &film.SpecialFeatures); err != nil {
			return nil, err
		}
		films = append(films, film)
	}

	return films, nil
}

func (r *filmPgRepository) GetById(id int) (*domain.Film, error) {
	ctx := context.Background()
	query := `SELECT film_id, title, description, release_year, language_id,
        original_language_id, rental_duration, rental_rate, length, replacement_cost,
        rating, last_update, special_features, fulltext
        FROM film WHERE film_id = $1`

	row := r.db.QueryRow(ctx, query, id)
	var film domain.Film
	var specialFeatures []byte

	err := row.Scan(
		&film.FilmID, &film.Title, &film.Description, &film.ReleaseYear,
		&film.LanguageID, &film.OriginalLangID, &film.RentalDuration,
		&film.RentalRate, &film.Length, &film.ReplacementCost,
		&film.Rating, &film.LastUpdate, &specialFeatures, &film.FullText,
	)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(specialFeatures, &film.SpecialFeatures); err != nil {
		return nil, err
	}

	return &film, nil
}

func (r *filmPgRepository) GetFilmsWithLanguage() ([]domain.FilmWithLanguage, error) {
	ctx := context.Background()

	query := `SELECT f.title, l.name
              FROM film f
              JOIN language l ON f.language_id = l.language_id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.FilmWithLanguage

	for rows.Next() {
		var row domain.FilmWithLanguage

		if err := rows.Scan(&row.Title, &row.LanguageName); err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, nil
}
