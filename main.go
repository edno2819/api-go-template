package main

import (
	"context"
	"gin-mongo-api/setting"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	setting.Setup()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(validateHeaders())

	productsRoutes := router.Group("/products")
	productsRoutes.Use(validateSecure)
	{
		productsRoutes.POST("/", createProduct)
		productsRoutes.GET("/", getProducts)
	}

	router.Run(":8080")
}

// Utilizar JWT token
// passar a instância do banco pelo midware
// Fazer o test de estress com o k6
// Criar middleware de erro
// fazer algum endpoint com processamento pesado
// criar endpoint com graphQL
// criar cache
