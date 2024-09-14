package utils_middleware

import (
	"gin-mongo-api/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDBFromContext(ctx *gin.Context) *mongo.Database {
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
