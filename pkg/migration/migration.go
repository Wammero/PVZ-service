package migration

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func callMigrations(dsn string) (*migrate.Migrate, error) {
	migrationsPath := "file://migrations"

	m, err := migrate.New(migrationsPath, dsn)

	if err != nil {
		return nil, err
	}

	return m, err
}

func ApplyMigrations(connStr string) {
	m, err := callMigrations(connStr)
	if err != nil {
		log.Fatalf("Error initializing migrations: %v", err)
	}

	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("Migrations are already applied")
		} else {
			log.Fatalf("Error applying migrations: %v", err)
		}
	} else {
		log.Println("Migrations successfully applied")
	}
}
