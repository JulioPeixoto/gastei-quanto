package auth

import (
	"errors"
	"sync"
)

type Repository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)
}

type memoryRepository struct {
	users map[string]*User
	mu    sync.RWMutex
}

func NewRepository() Repository {
	return &memoryRepository{
		users: make(map[string]*User),
	}
}

func (r *memoryRepository) Create(user *User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.Email == user.Email {
			return errors.New("email already exists")
		}
	}

	r.users[user.ID] = user
	return nil
}

func (r *memoryRepository) FindByEmail(email string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *memoryRepository) FindByID(id string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}




