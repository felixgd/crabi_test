package main

import (
	"crabi_test/app"
	"crabi_test/handlers"
	"crabi_test/repositories/mongodb"
	"crabi_test/repositories/pld"
	"crabi_test/service"
	"crabi_test/utils/constants"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router.
	r := gin.Default()

	pr := pld.PLD{}
	// Create the MongoDB client
	mr, err := mongodb.NewMongoClient(constants.MONGODB_URI)
	if err != nil {
		log.Fatal(err)
	}
	defer mr.Close()

	// Create a new instance of the MyService.
	s := service.NewUserService(pr, mr)

	// Create a new instance of the Handler and inject the service dependency.
	h := handlers.NewHandler(s)

	// Create a new instance of the Router and register the routes.
	router := app.NewRouter(r, h)
	router.RegisterRoutes()

	// Start the server.
	r.Run(":8080")
}
