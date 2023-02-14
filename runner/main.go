package main

import (
	"fmt"
	handler "github.com/margostino/openearth/api"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		message := fmt.Sprintf("Command Not Found!\n" +
			"Commands available: \n" +
			"- server: to start a local server\n" +
			"- schema-gen: to generate a new Graphql Schema")
		log.Panicln(message)
	}
	switch action := os.Args[1]; action {
	case "server":
		runLocalServer()
	case "schema-gen":
		//generateSchema()
	default:
		log.Printf("command not valid: %s\n", action)
	}
}

func runLocalServer() {
	http.HandleFunc("/query", handler.Query)
	http.HandleFunc("/playground", handler.Playground)
	log.Println("Starting OpenEarth-API Server in :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
