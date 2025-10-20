package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableToWatch, downCreateTableToWatch)
}

func upCreateTableToWatch(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`create table to_watch("user" uuid not null, "movie" uuid not null, created_at timestamp, updated_at timestamp, foreign key("user") references users(id), foreign key("movie") references movies(id), primary key("user", "movie"))`)
	return err
}

func downCreateTableToWatch(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`drop table to_watch`)
	return err
}
