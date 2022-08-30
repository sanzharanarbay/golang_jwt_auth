package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"jwt_auth_golang/application/routes"
)

func main() {
	port := os.Getenv("APP_PORT")

	prefix := os.Getenv("ROUTE_PREFIX")
	fmt.Println("Server started at " + port + "...")

	router := gin.New()
	// Routes
	routes.ApiRoutes(prefix, router)

	//Start Server on the port set in your .env file
	router.Run(":"+port)
}
