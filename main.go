package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yuuu/gqlgen-echo-sample/graph"
	"github.com/yuuu/gqlgen-echo-sample/graph/generated"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			"localhost", 5000, "postgres", "postgres", "password",
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", welcome())

		graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: db}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	er := e.Start(":1323")
	if er != nil {
		log.Fatalln(er)
	}
}

func welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	}
}