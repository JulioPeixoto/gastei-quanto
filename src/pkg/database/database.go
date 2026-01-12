package database

import (
	"database/sql"
)

type Database interface {
	GetDB() *sql.DB
	Close() error
	Migrate() error
}

type Config struct {
	Driver string
	DSN    string
}

