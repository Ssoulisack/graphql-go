package main

import (
	"fmt"
	"go-fiber/api/graph/resolver"
	"go-fiber/api/routes"
	"go-fiber/bootstrap"
	"log"
)


func main() {
	// Initialize the Fiber application
	app := bootstrap.App()
	globalEnv := app.Env
	fiber := app.Fiber
	db := app.DB

	// Setup routes for Fiber (including GraphQL)
	routes.Setup(fiber, db)

	// Call InitializeServices to get resolver
	resolver := resolver.InitializeServices(db)

	// Setup GraphQL routes
	routes.SetupGraphQL(fiber, resolver)

	log.Fatal(fiber.Listen(fmt.Sprintf(":%d", globalEnv.App.Port)))

}
