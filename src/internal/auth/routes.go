package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	auth := group.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)

		auth.POST("/logout", authMiddleware, handler.Logout)
		auth.GET("/me", authMiddleware, handler.Me)
	}
}
