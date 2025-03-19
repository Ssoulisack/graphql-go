package routes

import (
	"go-fiber/api/graph/resolver"
	"net/http"

	generated "go-fiber/api/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api", func(c *fiber.Ctx) error {
		return c.Next()
	})

	// Book routes
	BookRoutes(api, db)
}

// GraphQLRoutes sets up GraphQL routes for Fiber// GraphQLRoutes sets up the GraphQL routes
func SetupGraphQL(app *fiber.App, res *resolver.Resolver) {
	routes := app.Group("/api/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: res}))

	// Convert Fiber request to standard HTTP request for gqlgen
	routes.Post("/graphql", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(srv.ServeHTTP))(c.Context())
		return nil
	})
}
