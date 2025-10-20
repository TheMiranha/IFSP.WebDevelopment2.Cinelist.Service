package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableActors, downCreateTableActors)
}

func upCreateTableActors(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`create table actors(id uuid primary key, name varchar(255), imageUrl varchar(255), createdAt timestamp, updatedAt timestamp)`)
	return err
}

func downCreateTableActors(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`drop table actors`)
	return err
}
