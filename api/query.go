package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/margostino/openearth/graph"
	"net/http"
)

var server = newServer()

func Query(w http.ResponseWriter, r *http.Request) {
	//go metrics.PublishRequest(r)
	server.ServeHTTP(w, r)
}

func newServer() *handler.Server {
	c := graph.Config{Resolvers: &graph.Resolver{}}
	return handler.NewDefaultServer(graph.NewExecutableSchema(c))
}
