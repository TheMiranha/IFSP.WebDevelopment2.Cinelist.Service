package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableWatched, downCreateTableWatched)
}

func upCreateTableWatched(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`create table watched("user" uuid not null, "movie" uuid not null, rate numeric(10, 2), description varchar(255), created_at timestamp, updated_at timestamp, foreign key("user") references users(id), foreign key("movie") references movies(id), primary key("user", "movie"))`)
	return err
}

func downCreateTableWatched(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`drop table watched`)
	return err
}
