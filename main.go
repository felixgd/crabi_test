package main

import (
	"crabi_test/app"
	"crabi_test/handlers"
	"crabi_test/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router.
	r := gin.Default()

	// Create a new instance of the MyService.
	s := &service.MyService{}

	// Create a new instance of the Handler and inject the service dependency.
	h := handlers.NewHandler(s)

	// Create a new instance of the Router and register the routes.
	router := app.NewRouter(r, h)
	router.RegisterRoutes()

	// Start the server.
	r.Run(":8080")
}
