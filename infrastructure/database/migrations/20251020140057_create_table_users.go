package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableUsers, downCreateTableUsers)
}

func upCreateTableUsers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("create table users(id uuid primary key, email varchar(255), password varchar(255), name varchar(255), imageUrl varchar(255), createdAt timestamp, updatedAt timestamp)")
	return err
}

func downCreateTableUsers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("drop table users")
	return err
}
