package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up001, Down001)
}

// Up001 up migration.
func Up001(tx *sql.Tx) error {
	_, err := tx.Exec(
		`CREATE TABLE events (
			id              uuid primary key,
            user_id         uuid,
            title           text,
            date_and_time   timestamp with time zone,
            duration        int,
            description     text,
            time_to_notify  int
            )`)
	if err != nil {
		return err
	}

	return nil
}

// Down001 down migration.
func Down001(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE events;")
	if err != nil {
		return err
	}

	return nil
}
