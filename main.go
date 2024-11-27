package main

import (
	"log"
	"net/http"

	"github.com/ficoback/api"
)

func main() {
	handleGetScore()

	port := ":8080"

	log.Println("Starting server on " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleGetScore() {
	http.HandleFunc("/getScore", api.HandleGetScore)
}
