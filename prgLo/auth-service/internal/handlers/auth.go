package handlers

import (
	"github.com/LizaMeytner/prgLo/auth-service/internal/core"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, authCore *core.AuthCore) {
	r.POST("/register", func(c *gin.Context) {
		var request struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=8"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		// Используем контекст из Gin
		token, err := authCore.Register(c.Request.Context(), request.Email, request.Password)
		if err != nil {
			statusCode := 500
			if err == core.ErrUserExists {
				statusCode = 409
			}
			c.JSON(statusCode, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"status": "User created",
			"token":  token,
		})
	})
}
