package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDatabase struct {
	db *sql.DB
}

func NewSQLiteDatabase(dsn string) (Database, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &SQLiteDatabase{db: db}, nil
}

func (s *SQLiteDatabase) GetDB() *sql.DB {
	return s.db
}

func (s *SQLiteDatabase) Close() error {
	return s.db.Close()
}

func (s *SQLiteDatabase) Migrate() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS expenses (
			id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			date DATETIME NOT NULL,
			description TEXT NOT NULL,
			category TEXT,
			amount REAL NOT NULL,
			type TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)`,
		`CREATE INDEX IF NOT EXISTS idx_expenses_user_id ON expenses(user_id)`,
		`CREATE INDEX IF NOT EXISTS idx_expenses_date ON expenses(date)`,
		`CREATE INDEX IF NOT EXISTS idx_expenses_category ON expenses(category)`,
		`CREATE INDEX IF NOT EXISTS idx_expenses_type ON expenses(type)`,
	}

	for _, query := range queries {
		if _, err := s.db.Exec(query); err != nil {
			return err
		}
	}

	return nil
}
