package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Header 'X-API-KEY' é obrigatório"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func ValidateSecure(c *gin.Context) {
	apiKey := c.GetHeader("Authorization")
	if apiKey == "" || !strings.Contains(apiKey, "Bearer") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Autentificação é obrigatório"})
		c.Abort()
		return
	}
	token_array := strings.Split(apiKey, " ")
	if len(token_array) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Autentificação é obrigatório"})
		c.Abort()
		return
	}
	token := token_array[1]
	if token != "123456789" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Autentificação inválida!"})
		c.Abort()
		return
	}

	c.Next()
}
