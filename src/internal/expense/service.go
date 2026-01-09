package expense

import (
	"time"

	"github.com/google/uuid"
)

type Service interface {
	Create(userID string, req CreateExpenseRequest) (*Expense, error)
	GetByID(id, userID string) (*Expense, error)
	List(userID string, query ListExpensesQuery) ([]*Expense, error)
	Update(id, userID string, req UpdateExpenseRequest) (*Expense, error)
	Delete(id, userID string) error
	GetStats(userID string, startDate, endDate *time.Time) (*ExpenseStats, error)
	ImportTransactions(userID string, transactions []Transaction) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(userID string, req CreateExpenseRequest) (*Expense, error) {
	now := time.Now()

	expense := &Expense{
		ID:          uuid.New().String(),
		UserID:      userID,
		Date:        req.Date,
		Description: req.Description,
		Category:    req.Category,
		Amount:      req.Amount,
		Type:        req.Type,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.repo.Create(expense); err != nil {
		return nil, err
	}

	return expense, nil
}

func (s *service) GetByID(id, userID string) (*Expense, error) {
	return s.repo.FindByID(id, userID)
}

func (s *service) List(userID string, query ListExpensesQuery) ([]*Expense, error) {
	return s.repo.FindByUserID(userID, query)
}

func (s *service) Update(id, userID string, req UpdateExpenseRequest) (*Expense, error) {
	expense, err := s.repo.FindByID(id, userID)
	if err != nil {
		return nil, err
	}

	if req.Date != nil {
		expense.Date = *req.Date
	}

	if req.Description != nil {
		expense.Description = *req.Description
	}

	if req.Category != nil {
		expense.Category = *req.Category
	}

	if req.Amount != nil {
		expense.Amount = *req.Amount
	}

	if req.Type != nil {
		expense.Type = *req.Type
	}

	if err := s.repo.Update(expense); err != nil {
		return nil, err
	}

	return expense, nil
}

func (s *service) Delete(id, userID string) error {
	return s.repo.Delete(id, userID)
}

func (s *service) GetStats(userID string, startDate, endDate *time.Time) (*ExpenseStats, error) {
	return s.repo.GetStats(userID, startDate, endDate)
}

func (s *service) ImportTransactions(userID string, transactions []Transaction) (int, error) {
	count := 0

	for _, t := range transactions {
		expenseType := "expense"
		if t.Amount > 0 {
			expenseType = "income"
		}

		amount := t.Amount
		if amount < 0 {
			amount = -amount
		}

		expense := &Expense{
			ID:          uuid.New().String(),
			UserID:      userID,
			Date:        t.Date,
			Description: t.Description,
			Category:    t.Category,
			Amount:      amount,
			Type:        expenseType,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := s.repo.Create(expense); err != nil {
			return count, err
		}

		count++
	}

	return count, nil
}
