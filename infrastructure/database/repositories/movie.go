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

	movies := make([]entities.Movie, 0)
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

func (repo *MovieRepository) SearchByTitle(term string) ([]entities.Movie, error) {
	query := `SELECT id, title, description, image_url, released_at, created_at, updated_at, tmdb_rate 
	          FROM movies 
	          WHERE LOWER(title) LIKE LOWER($1)`

	searchTerm := "%" + term + "%"
	rows, err := repo.db.Query(query, searchTerm)

	if err != nil {
		return []entities.Movie{}, err
	}

	movies := make([]entities.Movie, 0)
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

	return movies, nil
}

func (repo *MovieRepository) GetCastByMovieID(movieID string) ([]entities.Actor, error) {
	query := `SELECT a.id, a.name, a.image_url, a.created_at, a.updated_at 
	          FROM actors a 
	          INNER JOIN casts c ON a.id = c.actor 
	          WHERE c.movie = $1`

	rows, err := repo.db.Query(query, movieID)

	if err != nil {
		return []entities.Actor{}, err
	}

	actors := make([]entities.Actor, 0)
	var actor entities.Actor

	for rows.Next() {
		err := rows.Scan(
			&actor.ID,
			&actor.Name,
			&actor.ImageUrl,
			&actor.CreatedAt,
			&actor.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
		} else {
			actors = append(actors, actor)
		}
	}

	return actors, nil
}
