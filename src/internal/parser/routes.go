package parser

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup, handler *Handler) {
	parser := group.Group("/parser")
	{
		parser.POST("/upload/csv", handler.UploadCSV)
		parser.POST("/import-and-save", handler.ImportAndSaveCSV)
	}
}
