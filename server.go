package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/wambugucoder/golang-graphql/config"
	"github.com/wambugucoder/golang-graphql/graph"
	"github.com/wambugucoder/golang-graphql/graph/generated"
	"log"
)

const defaultPort = ":8080"

func GraphQlHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
func PlayGroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func SetPort() string {
	//PORT
	port := config.ExtractKey("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}
func main() {
	app := gin.Default()
	//CONNECT DB
	config.Connect()

	app.POST("/query", GraphQlHandler())
	app.GET("/", PlayGroundHandler())
	err := app.Run(SetPort())
	if err != nil {
		log.Fatal("Server Could Not Start")
	}
	log.Fatal("Server Started on Port :" + SetPort())

}
