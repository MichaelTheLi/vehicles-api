package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"vehicles_api/api/graph"
	"vehicles_api/api/graph/resolvers"
	"vehicles_api/domain"
)

type ServerState struct {
	VehicleRepo domain.VehicleRepository
}

func BuildServer(state *ServerState, config Config) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// GraphQL Playground
	e.GET("/", func(c echo.Context) error {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// GraphQL endpoint
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolvers.NewResolver(
			state.VehicleRepo,
		),
	}))

	e.POST("/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	return e
}
