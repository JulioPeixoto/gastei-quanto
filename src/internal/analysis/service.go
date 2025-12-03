package analysis

import (
	"sort"
	"strings"
)

type Service interface {
	AnalyzeTransactions(transactions []Transaction) *AnalysisResponse
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) AnalyzeTransactions(transactions []Transaction) *AnalysisResponse {
	categoryMap := make(map[string]*CategorySummary)
	descriptionMap := make(map[string]*DescriptionSummary)

	var totalSpent, totalIncome float64

	for _, t := range transactions {
		if t.Amount > 0 {
			totalSpent += t.Amount
		} else {
			totalIncome += -t.Amount
		}

		category := t.Category
		if category == "" {
			category = inferCategory(t.Description)
		}

		if cat, exists := categoryMap[category]; exists {
			cat.Total += t.Amount
			cat.Count++
		} else {
			categoryMap[category] = &CategorySummary{
				Category: category,
				Total:    t.Amount,
				Count:    1,
			}
		}

		cleanDesc := cleanDescription(t.Description)
		if desc, exists := descriptionMap[cleanDesc]; exists {
			desc.Total += t.Amount
			desc.Count++
		} else {
			descriptionMap[cleanDesc] = &DescriptionSummary{
				Description: cleanDesc,
				Total:       t.Amount,
				Count:       1,
			}
		}
	}

	byCategory := make([]CategorySummary, 0, len(categoryMap))
	for _, cat := range categoryMap {
		if cat.Count > 0 {
			cat.Average = cat.Total / float64(cat.Count)
		}
		byCategory = append(byCategory, *cat)
	}
	sort.Slice(byCategory, func(i, j int) bool {
		return byCategory[i].Total > byCategory[j].Total
	})

	byDescription := make([]DescriptionSummary, 0, len(descriptionMap))
	for _, desc := range descriptionMap {
		byDescription = append(byDescription, *desc)
	}
	sort.Slice(byDescription, func(i, j int) bool {
		return byDescription[i].Total > byDescription[j].Total
	})

	return &AnalysisResponse{
		TotalSpent:       totalSpent,
		TotalIncome:      totalIncome,
		NetBalance:       totalIncome - totalSpent,
		TransactionCount: len(transactions),
		ByCategory:       byCategory,
		ByDescription:    byDescription,
	}
}

func inferCategory(description string) string {
	desc := strings.ToLower(description)

	if strings.Contains(desc, "99") || strings.Contains(desc, "ride") {
		return "Transporte"
	}
	if strings.Contains(desc, "ifood") || strings.Contains(desc, "pizza") ||
		strings.Contains(desc, "panif") || strings.Contains(desc, "tempero") {
		return "Alimentação"
	}
	if strings.Contains(desc, "amazon") || strings.Contains(desc, "mercado") {
		return "Compras"
	}
	if strings.Contains(desc, "spotify") || strings.Contains(desc, "prime") {
		return "Assinaturas"
	}
	if strings.Contains(desc, "iof") {
		return "Taxas"
	}
	if strings.Contains(desc, "pagamento recebido") || strings.Contains(desc, "estorno") || strings.Contains(desc, "crédito") {
		return "Receitas"
	}
	if strings.Contains(desc, "cursor") {
		return "Software"
	}

	return "Outros"
}

func cleanDescription(desc string) string {
	desc = strings.TrimSpace(desc)
	desc = strings.ReplaceAll(desc, "Pg *", "")
	desc = strings.ReplaceAll(desc, "Dl*", "")
	desc = strings.ReplaceAll(desc, "Dl *", "")
	desc = strings.ReplaceAll(desc, "Dm *", "")

	if strings.Contains(desc, " - Parcela") {
		idx := strings.Index(desc, " - Parcela")
		desc = desc[:idx]
	}

	return strings.TrimSpace(desc)
}
