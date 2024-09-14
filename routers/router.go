package routers

import (
	"gin-mongo-api/database"
	"gin-mongo-api/middleware"
	v1 "gin-mongo-api/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	db := database.ConnectDb()
	r.Use(middleware.DBMiddleware(db))
	r.Use(middleware.ValidateHeaders())

	apiv1 := r.Group("api/v1")
	apiv1.POST("/login", Login)
	apiv1.Use(middleware.JWTAuthMiddleware())
	{
		apiv1.POST("/person", v1.CreatePerson)
		apiv1.GET("/people", v1.GetPeople)
		apiv1.GET("/product", v1.GetProducts)
	}

	return r
}
