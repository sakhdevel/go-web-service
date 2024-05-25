package main

import (
	"github.com/gin-gonic/gin"
	"sakhdevel/go-web-service/db"
	"sakhdevel/go-web-service/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8085") // localhost
}
