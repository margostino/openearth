package main

import (
	handler "github.com/margostino/openearth/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/query", handler.Query)
	http.HandleFunc("/playground", handler.Playground)
	log.Println("Starting OpenEarth-API Server in :8080")
	log.Println("Playground: http://localhost:8080/playground")
	log.Println("GraphQL API: http://localhost:8080/query")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
