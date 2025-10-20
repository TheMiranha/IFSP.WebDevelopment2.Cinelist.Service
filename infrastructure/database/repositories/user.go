package repositories

import (
	"cinelist/domain/entities"
	"database/sql"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Create(user entities.User) (entities.User, error) {
	query, err := repo.db.Prepare("insert into users(id, email, password, name, createdAt, updatedAt) values ($1, $2, $3, $4, $5, $6)")

	if err != nil {
		return entities.User{}, err
	}

	query.QueryRow(user.ID, user.Email, user.Password, user.Name, time.Now(), time.Now())

	return user, nil
}
