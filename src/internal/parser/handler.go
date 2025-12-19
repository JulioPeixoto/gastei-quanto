package parser

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// UploadCSV godoc
// @Summary Upload CSV
// @Description Faz upload de um arquivo CSV contendo transações
// @Tags parser
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formData file true "CSV file"
// @Success 200 {object} parser.UploadResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /parser/upload/csv [post]
func (h *Handler) UploadCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Arquivo não encontrado. Use o campo 'file' no form-data",
		})
		return
	}

	if file.Header.Get("Content-Type") != "text/csv" &&
		!strings.HasSuffix(file.Filename, ".csv") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Apenas arquivos CSV são aceitos",
		})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao abrir arquivo",
		})
		return
	}
	defer f.Close()

	transactions, err := h.service.ParseCSV(f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao processar CSV: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UploadResponse{
		Message:      "CSV processado com sucesso",
		Count:        len(transactions),
		Transactions: transactions,
	})
}
