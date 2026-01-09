package expense

import (
	"errors"
	"sort"
	"sync"
	"time"
)

type Repository interface {
	Create(expense *Expense) error
	FindByID(id, userID string) (*Expense, error)
	FindByUserID(userID string, query ListExpensesQuery) ([]*Expense, error)
	Update(expense *Expense) error
	Delete(id, userID string) error
	GetStats(userID string, startDate, endDate *time.Time) (*ExpenseStats, error)
}

type memoryRepository struct {
	expenses map[string]*Expense
	mu       sync.RWMutex
}

func NewRepository() Repository {
	return &memoryRepository{
		expenses: make(map[string]*Expense),
	}
}

func (r *memoryRepository) Create(expense *Expense) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.expenses[expense.ID] = expense
	return nil
}

func (r *memoryRepository) FindByID(id, userID string) (*Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	expense, exists := r.expenses[id]
	if !exists {
		return nil, errors.New("expense not found")
	}

	if expense.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	return expense, nil
}

func (r *memoryRepository) FindByUserID(userID string, query ListExpensesQuery) ([]*Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*Expense

	for _, expense := range r.expenses {
		if expense.UserID != userID {
			continue
		}

		if !r.matchesQuery(expense, query) {
			continue
		}

		result = append(result, expense)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Date.After(result[j].Date)
	})

	return result, nil
}

func (r *memoryRepository) Update(expense *Expense) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, exists := r.expenses[expense.ID]
	if !exists {
		return errors.New("expense not found")
	}

	if existing.UserID != expense.UserID {
		return errors.New("unauthorized")
	}

	expense.UpdatedAt = time.Now()
	r.expenses[expense.ID] = expense
	return nil
}

func (r *memoryRepository) Delete(id, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	expense, exists := r.expenses[id]
	if !exists {
		return errors.New("expense not found")
	}

	if expense.UserID != userID {
		return errors.New("unauthorized")
	}

	delete(r.expenses, id)
	return nil
}

func (r *memoryRepository) GetStats(userID string, startDate, endDate *time.Time) (*ExpenseStats, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	stats := &ExpenseStats{}

	for _, expense := range r.expenses {
		if expense.UserID != userID {
			continue
		}

		if startDate != nil && expense.Date.Before(*startDate) {
			continue
		}

		if endDate != nil && expense.Date.After(*endDate) {
			continue
		}

		stats.Count++

		if expense.Type == "income" {
			stats.TotalIncome += expense.Amount
			stats.IncomeCount++
		} else {
			stats.TotalExpense += expense.Amount
			stats.ExpenseCount++
		}
	}

	stats.Balance = stats.TotalIncome - stats.TotalExpense

	return stats, nil
}

func (r *memoryRepository) matchesQuery(expense *Expense, query ListExpensesQuery) bool {
	if query.StartDate != nil && expense.Date.Before(*query.StartDate) {
		return false
	}

	if query.EndDate != nil && expense.Date.After(*query.EndDate) {
		return false
	}

	if query.Category != "" && expense.Category != query.Category {
		return false
	}

	if query.Type != "" && expense.Type != query.Type {
		return false
	}

	if query.MinAmount != nil && expense.Amount < *query.MinAmount {
		return false
	}

	if query.MaxAmount != nil && expense.Amount > *query.MaxAmount {
		return false
	}

	if query.Description != "" {
		found := false
		for i := 0; i < len(expense.Description)-len(query.Description)+1; i++ {
			match := true
			for j := 0; j < len(query.Description); j++ {
				if toLower(expense.Description[i+j]) != toLower(query.Description[j]) {
					match = false
					break
				}
			}
			if match {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}
