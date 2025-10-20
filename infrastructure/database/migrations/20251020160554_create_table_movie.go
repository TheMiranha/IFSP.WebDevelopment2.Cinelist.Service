package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableMovie, downCreateTableMovie)
}

func upCreateTableMovie(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("create table movies(id uuid primary key, title varchar(255), descriotion varchar(255), releasedAt timestamp, imageUrl varchar(255), tmdbRate decimal(10, 2), createdAt timestamp, updatedAt timestamp)")
	return err
}

func downCreateTableMovie(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("drop table movies")
	return err
}
