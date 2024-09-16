package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-mongo-api/service"
	utils_middleware "gin-mongo-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	Name string `json:"name" bson:"name" binding:"required"`
	Age  int    `json:"age" bson:"age"`
}

// ErrorResponse é a estrutura para respostas de erro.
// @Description Estrutura usada para respostas de erro na API.
type ErrorResponse struct {
	Error string `json:"error"`
}

// CreatePerson godoc
// @Summary      Cria uma nova pessoa
// @Description  Insere uma nova pessoa no banco de dados MongoDB
// @Tags         people
// @Accept       json
// @Produce      json
// @Param        person  body      Person  true  "Pessoa a ser criada"
// @Success      200     {object}  Person   "Pessoa criada com sucesso!"
// @Failure      400     {object}  ErrorResponse  "Invalid input"
// @Failure      500     {object}  ErrorResponse  "Erro ao inserir pessoa no MongoDB"
// @Router       /api/v1/people [post]
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := utils_middleware.GetDBFromGinSet(c)
	collection := db.Collection("people")

	_, err := collection.InsertOne(ctx, person)
	if err != nil {
		err_msg := fmt.Sprintf("Erro ao inserir pessoa no MongoDB. Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_msg})
		return
	}

	suc_msg := fmt.Sprintf("Pessoa criada com sucesso! %+v", person)
	c.JSON(http.StatusOK, gin.H{"status": suc_msg})
}

// GetPeople godoc
// @Summary      Retorna todas as pessoas
// @Description  Obtém uma lista de todas as pessoas do banco de dados MongoDB
// @Tags         people
// @Produce      json
// @Success      200  {array}   Person
// @Failure      500  {object}  ErrorResponse   "error: 'Erro ao buscar pessoas'"
// @Router       /api/v1/people [get]
func GetPeople(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cache := utils_middleware.GetCacheFromContext(c)

	result := service.GetDataCache(cache, "people-all")
	if result != "" {
		var cachedPeople []Person
		if err := json.Unmarshal([]byte(result), &cachedPeople); err == nil {
			c.JSON(http.StatusOK, gin.H{"CACHE": cachedPeople})
			return
		}
	}

	db := utils_middleware.GetDBFromGinSet(c)
	collection := db.Collection("people")

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

	peopleJSON, err := json.Marshal(people)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao serializar os dados para o cache"})
		return
	}

	err = cache.Set(ctx, "people-all", peopleJSON, 10*time.Minute).Err()
	if err != nil {
		fmt.Println("Error ao salvar dados no cache!")
	}

	c.JSON(http.StatusOK, people)
}

// CreateProduct godoc
// @Summary      Cria um novo produto
// @Description  Cria um produto (exemplo, a função está vazia)
// @Tags         products
// @Produce      json
// @Success      200  {object}  Person
// @Router       /api/v1/products [post]
func CreateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// SimulePanic godoc
// @Summary      Simula um panico da API
// @Description  Obtém uma lista de todos os produtos (exemplo, a função está vazia)
// @Tags         panic
// @Produce      json
// @Success      200  {object}  Person
// @Router       /api/v1/panic [get]
func SimulePanic(c *gin.Context) {
	cache := utils_middleware.GetCacheFromContext(c)

	result := service.GetDataCache(cache, "people-all")
	if result != "" {
		c.JSON(http.StatusOK, result)
		return
	}
	panic("Erro forçado: teste do CustomRecovery")
}
