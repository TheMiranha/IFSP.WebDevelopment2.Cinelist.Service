package repositories

import (
	"cinelist/domain/entities"
	"database/sql"
	"fmt"
)

type MovieRepository struct {
	db            *sql.DB
	moviesInCache []entities.Movie
}

func NewMovieRepository(db *sql.DB) MovieRepository {
	return MovieRepository{db: db, moviesInCache: []entities.Movie{}}
}

func (repo *MovieRepository) GetAll() ([]entities.Movie, error) {
	if len(repo.moviesInCache) > 0 {
		return repo.moviesInCache, nil
	}

	query := "select id, title, description, image_url, released_at, created_at, updated_at, tmdb_rate from movies"

	rows, err := repo.db.Query(query)

	if err != nil {
		return []entities.Movie{}, err
	}

	var movies []entities.Movie
	var movie entities.Movie

	for rows.Next() {
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.ImageUrl,
			&movie.ReleasedAt,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			&movie.TMDBRate,
		)

		if err != nil {
			fmt.Println(err)
		} else {
			movies = append(movies, movie)
		}
	}

	repo.moviesInCache = movies

	return movies, nil
}
