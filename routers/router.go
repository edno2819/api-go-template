package routers

import (
	"gin-mongo-api/middleware"
	v1 "gin-mongo-api/routers/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	// apiv1.Use(jwt.JWT())
	apiv1 := r.Group("api/v1")
	apiv1.Use(middleware.ValidateSecure)
	{
		apiv1.Use(v1.CreatePerson)
	}

	return r
}
