package v1

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	Name string `json:"name" bson:"name" binding:"required"`
	Age  int    `json:"age" bson:"age"`
}

func CreatePerson(c *gin.Context) {
	var person Person

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}
	if person.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name missing"})
		return
	}

	collection := client.Database(os.Getenv("MONGO_DB")).Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, person)
	if err != nil {
		err_msg := fmt.Sprintf("Erro ao inserir pessoa no MongoDB. Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_msg})
		return
	}

	suc_msg := fmt.Sprintf("Pessoa criada com sucesso! %+v", person)
	c.JSON(http.StatusOK, gin.H{"status": suc_msg})
}

func GetPeople(c *gin.Context) {
	collection := client.Database("testdb").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pessoas"})
		return
	}
	defer cursor.Close(ctx)

	var people []Person
	if err = cursor.All(ctx, &people); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao decodificar os dados"})
		return
	}

	c.JSON(http.StatusOK, people)
}

func CreateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
