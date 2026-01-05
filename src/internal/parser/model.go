package parser

import "time"

type Transaction struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
}

type UploadResponse struct {
	Message      string        `json:"message"`
	Count        int           `json:"count"`
	Transactions []Transaction `json:"transactions"`
}
