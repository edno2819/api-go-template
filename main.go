package main

import (
	"gin-mongo-api/routers"
	"gin-mongo-api/setting"
)

func init() {
	setting.Setup()
}

func main() {
	routersInit := routers.InitRouter()
	routersInit.Run(":8080")
}
