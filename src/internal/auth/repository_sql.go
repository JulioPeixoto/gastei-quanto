package auth

import (
	"database/sql"
	"errors"
)

type sqlRepository struct {
	db *sql.DB
}

// NewSQLRepository creates a Repository backed by the provided *sql.DB.
// The returned repository uses the sqlRepository implementation to perform user operations against that database handle.
func NewSQLRepository(db *sql.DB) Repository {
	return &sqlRepository{
		db: db,
	}
}

func (r *sqlRepository) Create(user *User) error {
	query := `INSERT INTO users (id, email, password, created_at) VALUES (?, ?, ?, ?)`

	_, err := r.db.Exec(query, user.ID, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.email" {
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