package parser

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	parser := router.Group("/parser")
	{
		parser.POST("/upload/csv", handler.UploadCSV)
	}
}
