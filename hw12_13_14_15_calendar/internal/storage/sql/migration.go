package sqlstorage

import (
	"database/sql"

	calendarconfig "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs/calendar_config"
	// need for work with migrations.
	_ "github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func Migrate(config *calendarconfig.Config) error {
	db, err := sql.Open("postgres", config.Dsn)
	if err != nil {
		panic(err)
	}

	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err = goose.Up(db, "./"); err != nil {
		return err
	}
	return nil
}
