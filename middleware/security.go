package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

const DbContextKey string = "db"

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

func DBMiddleware(db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(DbContextKey, db)
		c.Next()
	}
}
