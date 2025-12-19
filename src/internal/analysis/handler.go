package analysis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// AnalyzeTransactions godoc
// @Summary Analisa transações
// @Description Agrupa e analisa transações por categoria e descrição
// @Tags analysis
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body AnalysisRequest true "Lista de transações"
// @Success 200 {object} AnalysisResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /analysis/transactions [post]
func (h *Handler) AnalyzeTransactions(c *gin.Context) {
	var req AnalysisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Formato de requisição inválido: " + err.Error(),
		})
		return
	}

	if len(req.Transactions) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nenhuma transação fornecida",
		})
		return
	}

	result := h.service.AnalyzeTransactions(req.Transactions)
	c.JSON(http.StatusOK, result)
}
