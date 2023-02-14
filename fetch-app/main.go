package main

import (
	"fetch-app/server"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Fetch App Service
// @version         1.0
// @description     A fetch app service with 3 endpoint: fetch, aggregate, token verify

// @host      http://localhost:9000
// @BasePath  /api
func main() {
	r := gin.Default()

	server.RegisterAPIService(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":9000")
}