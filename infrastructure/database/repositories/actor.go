package repositories

import (
	"cinelist/domain/entities"
	"database/sql"
	"fmt"
)

type ActorRepository struct {
	db            *sql.DB
	actorsInCache []entities.Actor
}

func NewActorRepository(db *sql.DB) *ActorRepository {
	return &ActorRepository{
		db:            db,
		actorsInCache: []entities.Actor{},
	}
}

func (repo *ActorRepository) GetAll() ([]entities.Actor, error) {
	if len(repo.actorsInCache) > 0 {
		return repo.actorsInCache, nil
	}

	query := "select id, name, image_url, created_at, updated_at from actors"

	rows, err := repo.db.Query(query)

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

	repo.actorsInCache = actors

	return actors, nil
}
