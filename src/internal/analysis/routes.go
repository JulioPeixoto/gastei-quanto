package analysis

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup, handler *Handler) {
	analysis := group.Group("/analysis")
	{
		analysis.POST("/transactions", handler.AnalyzeTransactions)
	}
}
