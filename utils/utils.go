package utils_middleware

import (
	"gin-mongo-api/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDBFromGinSet(ctx *gin.Context) *mongo.Database {
	dbInterface, exists := ctx.Get(middleware.DbContextKey)
	if !exists {
		log.Fatalln("Banco de dados não encontrado!")
	}

	db, ok := dbInterface.(*mongo.Database)
	if !ok {
		log.Fatalln("Erro ao carregar o banco de dados!")
	}
	return db
}

func GetCacheFromContext(ctx *gin.Context) *redis.Client {
	cacheInterface, exists := ctx.Get(middleware.CacheContextKey)
	if !exists {
		log.Fatalln("Banco de dados não encontrado!")
	}

	cache, ok := cacheInterface.(*redis.Client)
	if !ok {
		log.Fatalln("Erro ao carregar o banco de dados!")
	}
	return cache
}
