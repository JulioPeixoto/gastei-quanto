package expense

import "time"

type Expense struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateExpenseRequest struct {
	Date        time.Time `json:"date" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount" binding:"required"`
	Type        string    `json:"type" binding:"required,oneof=income expense"`
}

type UpdateExpenseRequest struct {
	Date        *time.Time `json:"date"`
	Description *string    `json:"description"`
	Category    *string    `json:"category"`
	Amount      *float64   `json:"amount"`
	Type        *string    `json:"type" binding:"omitempty,oneof=income expense"`
}

type ListExpensesQuery struct {
	StartDate   *time.Time `form:"start_date"`
	EndDate     *time.Time `form:"end_date"`
	Category    string     `form:"category"`
	Type        string     `form:"type" binding:"omitempty,oneof=income expense"`
	MinAmount   *float64   `form:"min_amount"`
	MaxAmount   *float64   `form:"max_amount"`
	Description string     `form:"description"`
}

type ImportTransactionsRequest struct {
	Transactions []Transaction `json:"transactions" binding:"required"`
}

type Transaction struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
}

type ExpenseStats struct {
	TotalIncome   float64 `json:"total_income"`
	TotalExpense  float64 `json:"total_expense"`
	Balance       float64 `json:"balance"`
	Count         int     `json:"count"`
	IncomeCount   int     `json:"income_count"`
	ExpenseCount  int     `json:"expense_count"`
}

