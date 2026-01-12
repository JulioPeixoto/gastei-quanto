package auth

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

type sqlRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) Repository {
	return &sqlRepository{
		db: db,
	}
}

func (r *sqlRepository) Create(user *User) error {
	query := `INSERT INTO users (id, email, password, created_at) VALUES (?, ?, ?, ?)`
	
	_, err := r.db.Exec(query, user.ID, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return errors.New("email already exists")
		}
		return err
	}

	return nil
}

func (r *sqlRepository) FindByEmail(email string) (*User, error) {
	query := `SELECT id, email, password, created_at FROM users WHERE email = ?`

	user := &User{}
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *sqlRepository) FindByID(id string) (*User, error) {
	query := `SELECT id, email, password, created_at FROM users WHERE id = ?`

	user := &User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}
