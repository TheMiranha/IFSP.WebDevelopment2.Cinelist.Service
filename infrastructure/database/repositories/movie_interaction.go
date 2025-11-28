package repositories

import (
	"cinelist/domain/entities"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type MovieInteractionRepository struct {
	db *sql.DB
}

func NewMovieInteractionRepository(db *sql.DB) *MovieInteractionRepository {
	return &MovieInteractionRepository{db: db}
}

// Favorite methods
func (repo *MovieInteractionRepository) CreateFavorite(userID uuid.UUID, movieID uuid.UUID) error {
	query := `INSERT INTO favorites ("user", "movie", created_at, updated_at)
	          VALUES ($1, $2, $3, $4)
	          ON CONFLICT ("user", "movie")
	          DO UPDATE SET updated_at = $4`

	now := time.Now()
	_, err := repo.db.Exec(query, userID, movieID, now, now)
	return err
}

func (repo *MovieInteractionRepository) DeleteFavorite(userID uuid.UUID, movieID uuid.UUID) error {
	query := `DELETE FROM favorites WHERE "user" = $1 AND "movie" = $2`
	_, err := repo.db.Exec(query, userID, movieID)
	return err
}

func (repo *MovieInteractionRepository) GetFavoriteByUserAndMovie(userID uuid.UUID, movieID uuid.UUID) (entities.Favorite, error) {
	query := `SELECT "user", "movie", created_at, updated_at
	          FROM favorites
	          WHERE "user" = $1 AND "movie" = $2`

	var favorite entities.Favorite
	err := repo.db.QueryRow(query, userID, movieID).Scan(
		&favorite.User,
		&favorite.Movie,
		&favorite.CreatedAt,
		&favorite.UpdatedAt,
	)

	if err != nil {
		return entities.Favorite{}, err
	}

	return favorite, nil
}

// ToWatch methods
func (repo *MovieInteractionRepository) CreateToWatch(userID uuid.UUID, movieID uuid.UUID) error {
	query := `INSERT INTO to_watch ("user", "movie", created_at, updated_at)
	          VALUES ($1, $2, $3, $4)
	          ON CONFLICT ("user", "movie")
	          DO UPDATE SET updated_at = $4`

	now := time.Now()
	_, err := repo.db.Exec(query, userID, movieID, now, now)
	return err
}

func (repo *MovieInteractionRepository) DeleteToWatch(userID uuid.UUID, movieID uuid.UUID) error {
	query := `DELETE FROM to_watch WHERE "user" = $1 AND "movie" = $2`
	_, err := repo.db.Exec(query, userID, movieID)
	return err
}

func (repo *MovieInteractionRepository) GetToWatchByUserAndMovie(userID uuid.UUID, movieID uuid.UUID) (entities.ToWatch, error) {
	query := `SELECT "user", "movie", created_at, updated_at
	          FROM to_watch
	          WHERE "user" = $1 AND "movie" = $2`

	var toWatch entities.ToWatch
	err := repo.db.QueryRow(query, userID, movieID).Scan(
		&toWatch.User,
		&toWatch.Movie,
		&toWatch.CreatedAt,
		&toWatch.UpdatedAt,
	)

	if err != nil {
		return entities.ToWatch{}, err
	}

	return toWatch, nil
}

// Watched methods
func (repo *MovieInteractionRepository) CreateWatched(watched entities.Watched) error {
	query := `INSERT INTO watched ("user", "movie", rate, description, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6)
	          ON CONFLICT ("user", "movie")
	          DO UPDATE SET rate = $3, description = $4, updated_at = $6`

	now := time.Now()
	_, err := repo.db.Exec(query, watched.User, watched.Movie, watched.Rate, watched.Description, now, now)
	return err
}

func (repo *MovieInteractionRepository) UpdateWatched(watched entities.Watched) error {
	query := `UPDATE watched
	          SET rate = $3, description = $4, updated_at = $5
	          WHERE "user" = $1 AND "movie" = $2`

	now := time.Now()
	_, err := repo.db.Exec(query, watched.User, watched.Movie, watched.Rate, watched.Description, now)
	return err
}

func (repo *MovieInteractionRepository) GetWatchedByUserAndMovie(userID uuid.UUID, movieID uuid.UUID) (entities.Watched, error) {
	query := `SELECT "user", "movie", rate, description, created_at, updated_at
	          FROM watched
	          WHERE "user" = $1 AND "movie" = $2`

	var watched entities.Watched
	err := repo.db.QueryRow(query, userID, movieID).Scan(
		&watched.User,
		&watched.Movie,
		&watched.Rate,
		&watched.Description,
		&watched.CreatedAt,
		&watched.UpdatedAt,
	)

	if err != nil {
		return entities.Watched{}, err
	}

	return watched, nil
}

func (repo *MovieInteractionRepository) GetFavoritesByUserID(userID uuid.UUID) ([]entities.Movie, error) {
	query := `SELECT m.id, m.title, m.description, m.image_url, m.released_at, m.created_at, m.updated_at, m.tmdb_rate
	          FROM movies m
	          INNER JOIN favorites f ON m.id = f.movie
	          WHERE f."user" = $1`

	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return []entities.Movie{}, err
	}

	movies := make([]entities.Movie, 0)
	var movie entities.Movie
	var tmdbRate sql.NullFloat64

	for rows.Next() {
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.ImageUrl,
			&movie.ReleasedAt,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			&tmdbRate,
		)

		if err != nil {
			continue
		}
		if tmdbRate.Valid {
			movie.TMDBRate = tmdbRate.Float64
		} else {
			movie.TMDBRate = 0.0
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (repo *MovieInteractionRepository) GetToWatchByUserID(userID uuid.UUID) ([]entities.Movie, error) {
	query := `SELECT m.id, m.title, m.description, m.image_url, m.released_at, m.created_at, m.updated_at, m.tmdb_rate
	          FROM movies m
	          INNER JOIN to_watch tw ON m.id = tw.movie
	          WHERE tw."user" = $1`

	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return []entities.Movie{}, err
	}

	movies := make([]entities.Movie, 0)
	var movie entities.Movie
	var tmdbRate sql.NullFloat64

	for rows.Next() {
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.ImageUrl,
			&movie.ReleasedAt,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			&tmdbRate,
		)

		if err != nil {
			continue
		}
		if tmdbRate.Valid {
			movie.TMDBRate = tmdbRate.Float64
		} else {
			movie.TMDBRate = 0.0
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (repo *MovieInteractionRepository) GetWatchedByUserID(userID uuid.UUID) ([]entities.Watched, error) {
	query := `SELECT "user", "movie", rate, description, created_at, updated_at
	          FROM watched
	          WHERE "user" = $1`

	rows, err := repo.db.Query(query, userID)
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
			continue
		}
		watchedList = append(watchedList, watched)
	}

	return watchedList, nil
}
