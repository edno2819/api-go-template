package middleware

import (
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

const DbContextKey string = "db"
const CacheContextKey string = "cache"

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

func CustomRecovery(c *gin.Context, recovered interface{}) {
	err, ok := recovered.(string)
	if !ok {
		err = "Erro desconhecido"
	}

	// Captura a stack trace
	stack := make([]byte, 1024*8)
	stack = stack[:runtime.Stack(stack, false)]

	log.Printf("Panic recuperado: %s\n%s", err, stack)

	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"message": "Erro interno do servidor",
	})

	c.AbortWithStatus(http.StatusInternalServerError)
}

func CacheMiddleware(cache *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(CacheContextKey, cache)
		c.Next()
	}
}
