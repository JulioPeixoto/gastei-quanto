package main

import (
	"gastei-quanto/src/internal/parser"
	"log"

	"github.com/gin-gonic/gin"

	_ "gastei-quanto/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gastei Quanto API
// @version 1.0
// @description API para análise de faturas do Nubank

// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		parserService := parser.NewService()
		parserHandler := parser.NewHandler(parserService)
		parser.RegisterRoutes(api, parserHandler)
	}

	log.Println("🚀 Servidor rodando na porta 8080")
	log.Println("📚 Swagger: http://localhost:8080/swagger/index.html")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
