package main

import (
	"fetch-app/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	server.RegisterAPIService(r)

	r.Run(":9000")
}