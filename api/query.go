package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/margostino/openearth/graph"
	"github.com/margostino/openearth/graph/generated"
	"net/http"
)

var server = newServer()

func Query(w http.ResponseWriter, r *http.Request) {
	//go metrics.PublishRequest(r)
	server.ServeHTTP(w, r)
}

func newServer() *handler.Server {
	c := generated.Config{Resolvers: &graph.Resolver{}}
	return handler.NewDefaultServer(generated.NewExecutableSchema(c))
}
