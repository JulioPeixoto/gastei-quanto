package expense

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
	expenses := rg.Group("/expenses")
	{
		expenses.POST("", handler.Create)
		expenses.GET("", handler.List)
		expenses.GET("/stats", handler.GetStats)
		expenses.GET("/:id", handler.GetByID)
		expenses.PUT("/:id", handler.Update)
		expenses.DELETE("/:id", handler.Delete)
		expenses.POST("/import", handler.ImportTransactions)
	}
}
