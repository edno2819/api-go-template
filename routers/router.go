package routers

import (
	"gin-mongo-api/database"
	"gin-mongo-api/middleware"
	v1 "gin-mongo-api/routers/api"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(middleware.CustomRecovery))

	db := database.ConnectDb()
	cache := database.ConnectRedis()

	r.Use(middleware.DBMiddleware(db))

	r.StaticFile("/swagger.yaml", "./docs/swagger.yaml")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.yaml")))

	r.Use(middleware.ValidateHeaders())
	apiv1 := r.Group("api/v1")
	apiv1.POST("/login", middleware.CacheMiddleware(cache), v1.Login)
	apiv1.Use(middleware.JWTAuthMiddleware())
	{
		apiv1.POST("/person", v1.CreatePerson)
		apiv1.GET("/people", middleware.CacheMiddleware(cache), v1.GetPeople)
		apiv1.GET("/panic", middleware.CacheMiddleware(cache), v1.SimulePanic)
	}

	return r
}
