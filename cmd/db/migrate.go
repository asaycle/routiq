package db

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		dbURL, ok := os.LookupEnv("DATABASE_URL")
		if !ok {
			log.Printf("Error env: %v\n", os.Environ())
			return nil
		}
		m, err := migrate.New(
			"file://db/migration",
			dbURL,
		)
		if err != nil {
			return xerrors.Errorf("Error creating migration instance: %w", err)
		}
		if err := m.Up(); err != nil {
			switch err {
			case migrate.ErrNoChange:
				return nil
			default:
				return xerrors.Errorf("Error running migrations: %w", err)
			}
		}
		return nil
	},
}
