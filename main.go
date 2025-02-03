package main

import (
	"log"
	"net/http"
	"Number-Classification-API/handlers"
)

func main () {
	http.HandleFunc("/api/classify-number", handlers.ClassifyNumberHandler)
	log.Println("Server started at http://localhost:8080/api/classify-number")
	log.Fatal(http.ListenAndServe(":8080", nil))
}