package main

import (
	"gastei-quanto/src/internal/analysis"
	"gastei-quanto/src/internal/auth"
	"gastei-quanto/src/internal/expense"
	"gastei-quanto/src/internal/parser"
	"log"
	"os"

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

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token
func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key-change-in-production"
	}

	authRepo := auth.NewRepository()
	authService := auth.NewService(authRepo, jwtSecret)
	authHandler := auth.NewHandler(authService)
	authMiddleware := auth.AuthMiddleware(authService)

	api := router.Group("/api/v1")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		auth.RegisterRoutes(api, authHandler, authMiddleware)

		protected := api.Group("")
		protected.Use(authMiddleware)
		{
			parserService := parser.NewService()
			parserHandler := parser.NewHandler(parserService)
			parser.RegisterRoutes(protected, parserHandler)

			analysisService := analysis.NewService()
			analysisHandler := analysis.NewHandler(analysisService)
			analysis.RegisterRoutes(protected, analysisHandler)

			expenseRepo := expense.NewRepository()
			expenseService := expense.NewService(expenseRepo)
			expenseHandler := expense.NewHandler(expenseService)
			expense.RegisterRoutes(protected, expenseHandler)
		}
	}

	log.Println("🚀 Servidor rodando na porta 8080")
	log.Println("📚 Swagger: http://localhost:8080/swagger/index.html")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
