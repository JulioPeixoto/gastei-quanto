package parser

import (
	"fmt"
	"gastei-quanto/src/internal/analysis"
	"gastei-quanto/src/internal/expense"
	"io"
	"log"
	"strings"
)

type IntegrationService interface {
	ProcessAndSaveCSV(userID string, file io.Reader) (*ImportAndSaveResponse, error)
}

type integrationService struct {
	parserService   Service
	analysisService analysis.Service
	expenseService  expense.Service
}

func NewIntegrationService(
	parserService Service,
	analysisService analysis.Service,
	expenseService expense.Service,
) IntegrationService {
	return &integrationService{
		parserService:   parserService,
		analysisService: analysisService,
		expenseService:  expenseService,
	}
}

func (s *integrationService) ProcessAndSaveCSV(userID string, file io.Reader) (*ImportAndSaveResponse, error) {
	if userID == "" {
		return nil, fmt.Errorf("userID não pode ser vazio")
	}

	if file == nil {
		return nil, fmt.Errorf("arquivo não pode ser nulo")
	}

	transactions, err := s.parserService.ParseCSV(file)
	if err != nil {
		return nil, fmt.Errorf("erro ao processar CSV: %w", err)
	}

	if len(transactions) == 0 {
		return nil, fmt.Errorf("nenhuma transação encontrada no arquivo CSV")
	}

	log.Printf("Parsed %d transactions from CSV for user %s", len(transactions), userID)

	categorizedTransactions := s.categorizeTransactions(transactions)

	expenseTransactions := s.convertToExpenseTransactions(categorizedTransactions)

	savedCount, err := s.expenseService.ImportTransactions(userID, expenseTransactions)
	if err != nil {
		log.Printf("Error saving transactions for user %s: %v", userID, err)
		return nil, fmt.Errorf("erro ao salvar transações: %w", err)
	}

	log.Printf("Successfully saved %d/%d transactions for user %s", savedCount, len(transactions), userID)

	return &ImportAndSaveResponse{
		Message:      "CSV processado e salvo com sucesso",
		Processed:    len(transactions),
		Saved:        savedCount,
		Transactions: categorizedTransactions,
	}, nil
}

func (s *integrationService) categorizeTransactions(transactions []Transaction) []Transaction {
	log.Printf("Starting categorization of %d transactions", len(transactions))
	
	analysisTransactions := make([]analysis.Transaction, len(transactions))
	for i, t := range transactions {
		analysisTransactions[i] = analysis.Transaction{
			Date:        t.Date.Format("2006-01-02"),
			Description: t.Description,
			Category:    t.Category,
			Amount:      t.Amount,
		}
	}

	result := s.analysisService.AnalyzeTransactions(analysisTransactions)
	log.Printf("Analysis completed: Total spent: %.2f, Total income: %.2f", result.TotalSpent, result.TotalIncome)

	categoryMap := make(map[string]string)
	for _, cat := range result.ByCategory {
		categoryMap[cat.Category] = cat.Category
	}

	categorizedCount := 0
	for i := range transactions {
		if transactions[i].Category == "" {
			category := s.suggestCategory(transactions[i].Description)
			transactions[i].Category = category
			categorizedCount++
			log.Printf("Auto-categorized [%d/%d] '%s' as '%s'", i+1, len(transactions), transactions[i].Description, category)
		}
	}

	log.Printf("Categorization complete: %d transactions auto-categorized", categorizedCount)

	return transactions
}

func (s *integrationService) suggestCategory(description string) string {
	descLower := strings.ToLower(description)

	if strings.Contains(descLower, "estorno") || strings.Contains(descLower, "crédito de") || strings.Contains(descLower, "credito de") || strings.Contains(descLower, "pagamento recebido") {
		return "Credito"
	}

	transportKeywords := []string{"uber", "99", "taxi", "ride", "dl*", "pg *", "dl *", "transporte", "estacionamento"}
	for _, keyword := range transportKeywords {
		if strings.Contains(descLower, keyword) {
			return "Transporte"
		}
	}

	foodKeywords := []string{"ifood", "restaurante", "padaria", "panif", "pizza", "lanche", "acai", "açai", "food", "bar", "cafe", "café", "tempero"}
	for _, keyword := range foodKeywords {
		if strings.Contains(descLower, keyword) {
			return "Alimentacao"
		}
	}

	shoppingKeywords := []string{"amazon", "mercado", "compra", "loja", "mercadolivre"}
	for _, keyword := range shoppingKeywords {
		if strings.Contains(descLower, keyword) {
			return "Compras"
		}
	}

	subscriptionKeywords := []string{"spotify", "netflix", "prime", "assinatura", "dm *"}
	for _, keyword := range subscriptionKeywords {
		if strings.Contains(descLower, keyword) {
			return "Assinaturas"
		}
	}

	if strings.Contains(descLower, "iof") {
		return "Taxas"
	}

	return "Outros"
}

func (s *integrationService) convertToExpenseTransactions(transactions []Transaction) []expense.Transaction {
	result := make([]expense.Transaction, len(transactions))
	for i, t := range transactions {
		result[i] = expense.Transaction{
			Date:        t.Date,
			Description: t.Description,
			Category:    t.Category,
			Amount:      t.Amount,
		}
	}
	return result
}

