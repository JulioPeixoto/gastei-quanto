package expense

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type sqlRepository struct {
	db *sql.DB
}

// NewSQLRepository constructs a Repository backed by the provided *sql.DB.
// The returned Repository uses the given database handle for SQL-based expense persistence.
func NewSQLRepository(db *sql.DB) Repository {
	return &sqlRepository{
		db: db,
	}
}

func (r *sqlRepository) Create(expense *Expense) error {
	query := `INSERT INTO expenses (id, user_id, date, description, category, amount, type, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(
		query,
		expense.ID,
		expense.UserID,
		expense.Date,
		expense.Description,
		expense.Category,
		expense.Amount,
		expense.Type,
		expense.CreatedAt,
		expense.UpdatedAt,
	)

	return err
}

func (r *sqlRepository) FindByID(id, userID string) (*Expense, error) {
	query := `SELECT id, user_id, date, description, category, amount, type, created_at, updated_at 
		FROM expenses WHERE id = ? AND user_id = ?`

	expense := &Expense{}
	err := r.db.QueryRow(query, id, userID).Scan(
		&expense.ID,
		&expense.UserID,
		&expense.Date,
		&expense.Description,
		&expense.Category,
		&expense.Amount,
		&expense.Type,
		&expense.CreatedAt,
		&expense.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("expense not found")
		}
		return nil, err
	}

	return expense, nil
}

func (r *sqlRepository) FindByUserID(userID string, query ListExpensesQuery) ([]*Expense, error) {
	queryStr := `SELECT id, user_id, date, description, category, amount, type, created_at, updated_at 
		FROM expenses WHERE user_id = ?`

	args := []interface{}{userID}
	conditions := []string{}

	if query.StartDate != nil {
		conditions = append(conditions, "date >= ?")
		args = append(args, query.StartDate)
	}

	if query.EndDate != nil {
		conditions = append(conditions, "date <= ?")
		args = append(args, query.EndDate)
	}

	if query.Category != "" {
		conditions = append(conditions, "category = ?")
		args = append(args, query.Category)
	}

	if query.Type != "" {
		conditions = append(conditions, "type = ?")
		args = append(args, query.Type)
	}

	if query.MinAmount != nil {
		conditions = append(conditions, "amount >= ?")
		args = append(args, query.MinAmount)
	}

	if query.MaxAmount != nil {
		conditions = append(conditions, "amount <= ?")
		args = append(args, query.MaxAmount)
	}

	if query.Description != "" {
		conditions = append(conditions, "description LIKE ?")
		args = append(args, "%"+query.Description+"%")
	}

	if len(conditions) > 0 {
		queryStr += " AND " + strings.Join(conditions, " AND ")
	}

	queryStr += " ORDER BY date DESC"

	rows, err := r.db.Query(queryStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	expenses := []*Expense{}
	for rows.Next() {
		expense := &Expense{}
		err := rows.Scan(
			&expense.ID,
			&expense.UserID,
			&expense.Date,
			&expense.Description,
			&expense.Category,
			&expense.Amount,
			&expense.Type,
			&expense.CreatedAt,
			&expense.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (r *sqlRepository) Update(expense *Expense) error {
	query := `UPDATE expenses SET date = ?, description = ?, category = ?, amount = ?, type = ?, updated_at = ? 
		WHERE id = ? AND user_id = ?`

	result, err := r.db.Exec(
		query,
		expense.Date,
		expense.Description,
		expense.Category,
		expense.Amount,
		expense.Type,
		expense.UpdatedAt,
		expense.ID,
		expense.UserID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("expense not found")
	}

	return nil
}

func (r *sqlRepository) Delete(id, userID string) error {
	query := `DELETE FROM expenses WHERE id = ? AND user_id = ?`

	result, err := r.db.Exec(query, id, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("expense not found")
	}

	return nil
}

func (r *sqlRepository) GetStats(userID string, startDate, endDate *time.Time) (*ExpenseStats, error) {
	query := `SELECT 
		COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0) as total_income,
		COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0) as total_expense,
		COUNT(*) as count,
		COALESCE(SUM(CASE WHEN type = 'income' THEN 1 ELSE 0 END), 0) as income_count,
		COALESCE(SUM(CASE WHEN type = 'expense' THEN 1 ELSE 0 END), 0) as expense_count
		FROM expenses WHERE user_id = ?`

	args := []interface{}{userID}

	if startDate != nil {
		query += " AND date >= ?"
		args = append(args, startDate)
	}

	if endDate != nil {
		query += " AND date <= ?"
		args = append(args, endDate)
	}

	stats := &ExpenseStats{}
	err := r.db.QueryRow(query, args...).Scan(
		&stats.TotalIncome,
		&stats.TotalExpense,
		&stats.Count,
		&stats.IncomeCount,
		&stats.ExpenseCount,
	)

	if err != nil {
		return nil, err
	}

	stats.Balance = stats.TotalIncome - stats.TotalExpense

	return stats, nil
}

var _ = fmt.Sprint("")