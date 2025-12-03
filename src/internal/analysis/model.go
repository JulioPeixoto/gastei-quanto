package analysis

type CategorySummary struct {
	Category string  `json:"category"`
	Total    float64 `json:"total"`
	Count    int     `json:"count"`
	Average  float64 `json:"average"`
}

type DescriptionSummary struct {
	Description string  `json:"description"`
	Total       float64 `json:"total"`
	Count       int     `json:"count"`
}

type AnalysisResponse struct {
	TotalSpent       float64              `json:"total_spent"`
	TotalIncome      float64              `json:"total_income"`
	NetBalance       float64              `json:"net_balance"`
	TransactionCount int                  `json:"transaction_count"`
	ByCategory       []CategorySummary    `json:"by_category"`
	ByDescription    []DescriptionSummary `json:"by_description"`
}

type AnalysisRequest struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Amount      float64 `json:"amount"`
}
