package database

import (
	"context"
	"database/sql"
	"embed"
	_ "embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed migrations/*.sql
var migrations embed.FS

type SQLite struct {
	Conn *sql.DB
}

func NewMemory() (*SQLite, error) {
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	if err := Migrate(conn); err != nil {
		return nil, err
	}

	return &SQLite{Conn: conn}, nil
}

func NewLocal(pathToDatabase string) (*SQLite, error) {
	uri := fmt.Sprintf("sqlite3://%s?journaled_mode=WAL", pathToDatabase)

	conn, err := sql.Open("sqlite3", uri)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	if err := Migrate(conn); err != nil {
		return nil, err
	}

	return &SQLite{Conn: conn}, nil
}

func Migrate(conn *sql.DB) error {
	// Prepare resources
	migrationSource, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}
	driver, err := sqlite.WithInstance(conn, &sqlite.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("iofs", migrationSource, "sqlite", driver)
	if err != nil {
		return err
	}

	// Migrate
	if err := m.Up(); err != nil {
		return err
	}

	// Free resources
	err1, err2 := m.Close()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func (db *SQLite) Queries() (*Queries, error) {
	return Prepare(context.Background(), db.Conn)
}
