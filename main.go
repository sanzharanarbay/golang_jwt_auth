package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"jwt_auth_golang/routes"
)

func main() {
	port := os.Getenv("APP_PORT")

	prefix := os.Getenv("ROUTE_PREFIX")
	fmt.Println("Server started at " + port + "...")
	r := mux.NewRouter().StrictSlash(true)

	// Routes
	routes.ApiRoutes(prefix, r)

	//Start Server on the port set in your .env file
	err := http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r))
	if err != nil {
		log.Fatal(err)
	}
}
