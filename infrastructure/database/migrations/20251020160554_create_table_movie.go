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
	_, err := tx.Exec("create table movies(id uuid primary key, title varchar(255), description varchar(255), released_at timestamp, image_url varchar(255), tmdb_rate decimal(10, 2), created_at timestamp, updated_at timestamp)")
	return err
}

func downCreateTableMovie(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("drop table movies")
	return err
}
