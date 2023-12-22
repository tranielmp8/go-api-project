package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tranielmp8/go-api-project/db"
	"github.com/tranielmp8/go-api-project/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080

}
