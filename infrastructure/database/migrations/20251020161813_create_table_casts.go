package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableCasts, downCreateTableCasts)
}

func upCreateTableCasts(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`create table casts("actor" uuid not null, "movie" uuid not null, createdAt timestamp, updatedAt timestamp, foreign key("actor") references actors(id), foreign key("movie") references movies(id), primary key("actor", "movie"))`)
	return err
}

func downCreateTableCasts(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`drop table casts`)
	return err
}
