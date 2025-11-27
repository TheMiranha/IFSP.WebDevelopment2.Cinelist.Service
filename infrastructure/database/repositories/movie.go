package repositories

import (
	"cinelist/domain/entities"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
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

func (repo *MovieRepository) GetClassics() ([]entities.Movie, error) {
	query := `SELECT id, title, description, image_url, released_at, created_at, updated_at, tmdb_rate 
	          FROM movies 
	          WHERE released_at IS NOT NULL 
	          ORDER BY released_at 
	          LIMIT 50`

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

	return movies, nil
}

func (repo *MovieRepository) GetHighlights() ([]entities.Movie, error) {
	query := `SELECT id, title, description, image_url, released_at, created_at, updated_at, tmdb_rate 
	          FROM movies 
	          WHERE tmdb_rate IS NOT NULL 
	          ORDER BY tmdb_rate DESC 
	          LIMIT 51`

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

	return movies, nil
}

func (repo *MovieRepository) GetNewReleases() ([]entities.Movie, error) {
	query := `SELECT id, title, description, image_url, released_at, created_at, updated_at, tmdb_rate 
	          FROM movies 
	          WHERE released_at IS NOT NULL 
	          ORDER BY released_at DESC 
	          LIMIT 50`

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

	return movies, nil
}

func (repo *MovieRepository) GetById(id string) (entities.Movie, error) {
	query := `SELECT id, title, description, image_url, released_at, created_at, updated_at, tmdb_rate 
	          FROM movies 
	          WHERE id = $1`

	var movie entities.Movie
	err := repo.db.QueryRow(query, id).Scan(
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
		return entities.Movie{}, err
	}

	return movie, nil
}

type RatingWithUser struct {
	UserID      uuid.UUID
	UserName    string
	UserImageUrl string
	Rate        float64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (repo *MovieRepository) GetWatchedByMovieID(movieID string) ([]entities.Watched, error) {
	query := `SELECT "user", "movie", rate, description, created_at, updated_at 
	          FROM watched 
	          WHERE "movie" = $1`

	rows, err := repo.db.Query(query, movieID)

	if err != nil {
		return []entities.Watched{}, err
	}

	watchedList := make([]entities.Watched, 0)
	var watched entities.Watched

	for rows.Next() {
		err := rows.Scan(
			&watched.User,
			&watched.Movie,
			&watched.Rate,
			&watched.Description,
			&watched.CreatedAt,
			&watched.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
		} else {
			watchedList = append(watchedList, watched)
		}
	}

	return watchedList, nil
}

func (repo *MovieRepository) GetRatingsWithUserByMovieID(movieID string) ([]RatingWithUser, error) {
	query := `SELECT w."user", u.name, u.image_url, w.rate, w.description, w.created_at, w.updated_at
	          FROM watched w
	          INNER JOIN users u ON w."user" = u.id
	          WHERE w."movie" = $1
	          ORDER BY w.created_at DESC`

	rows, err := repo.db.Query(query, movieID)

	if err != nil {
		return []RatingWithUser{}, err
	}

	ratings := make([]RatingWithUser, 0)
	var rating RatingWithUser

	for rows.Next() {
		err := rows.Scan(
			&rating.UserID,
			&rating.UserName,
			&rating.UserImageUrl,
			&rating.Rate,
			&rating.Description,
			&rating.CreatedAt,
			&rating.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
		} else {
			ratings = append(ratings, rating)
		}
	}

	return ratings, nil
}
