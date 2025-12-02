package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type Service interface {
	ParseCSV(file io.Reader) ([]Transaction, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) ParseCSV(file io.Reader) ([]Transaction, error) {
	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.TrimLeadingSpace = true

	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler header: %w", err)
	}

	dateIdx := findColumn(header, "date", "data")
	categoryIdx := findColumn(header, "category", "categoria")
	descIdx := findColumn(header, "title", "description", "titulo", "título", "descricao", "descrição")
	amountIdx := findColumn(header, "amount", "value", "valor")

	if dateIdx == -1 || amountIdx == -1 {
		return nil, fmt.Errorf("colunas obrigatórias não encontradas (date, amount)")
	}

	var transactions []Transaction

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		if len(record) <= dateIdx || len(record) <= amountIdx {
			continue
		}

		date, err := parseDate(record[dateIdx])
		if err != nil {
			continue
		}

		amount, err := parseAmount(record[amountIdx])
		if err != nil {
			continue
		}

		category := ""
		if categoryIdx >= 0 && len(record) > categoryIdx {
			category = record[categoryIdx]
		}

		description := ""
		if descIdx >= 0 && len(record) > descIdx {
			description = record[descIdx]
		}

		transactions = append(transactions, Transaction{
			Date:        date,
			Category:    category,
			Description: description,
			Amount:      amount,
		})
	}

	return transactions, nil
}

func findColumn(header []string, names ...string) int {
	for i, col := range header {
		colLower := strings.ToLower(strings.TrimSpace(col))
		for _, name := range names {
			if colLower == strings.ToLower(name) {
				return i
			}
		}
	}
	return -1
}

func parseDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)

	formats := []string{
		"2006-01-02",
		"02/01/2006",
		"01/02/2006",
		"2006/01/02",
	}

	for _, format := range formats {
		if date, err := time.Parse(format, dateStr); err == nil {
			return date, nil
		}
	}

	return time.Time{}, fmt.Errorf("formato de data inválido: %s", dateStr)
}

func parseAmount(amountStr string) (float64, error) {
	amountStr = strings.TrimSpace(amountStr)
	amountStr = strings.ReplaceAll(amountStr, "R$", "")
	amountStr = strings.ReplaceAll(amountStr, " ", "")
	amountStr = strings.ReplaceAll(amountStr, ",", ".")

	return strconv.ParseFloat(amountStr, 64)
}
