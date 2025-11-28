package repositories

import (
	"cinelist/domain/entities"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Create(user entities.User) (entities.User, error) {
	query, err := repo.db.Prepare("insert into users(id, email, password, name, created_at, updated_at) values ($1, $2, $3, $4, $5, $6)")

	if err != nil {
		return entities.User{}, err
	}

	query.QueryRow(user.ID, user.Email, user.Password, user.Name, time.Now(), time.Now())

	return user, nil
}

func (repo *UserRepository) GetById(id uuid.UUID) (entities.User, error) {
	query, err := repo.db.Prepare("select id, email, password, name, created_at, updated_at from users where id = $1")
	if err != nil {
		return entities.User{}, err
	}

	var user entities.User

	err = query.QueryRow(id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (entities.User, error) {
	query, err := repo.db.Prepare("select id, email, password, name, created_at, updated_at from users where email = $1")
	if err != nil {
		return entities.User{}, err
	}

	var user entities.User

	err = query.QueryRow(email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
