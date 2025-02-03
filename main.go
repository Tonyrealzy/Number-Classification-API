package main

import (
	"Number-Classification-API/handlers"
	"Number-Classification-API/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/api/classify-number", middleware.HandleCORS(), handlers.ClassifyNumberHandler)
	router.Run(":8080")
	log.Println("Server started at http://localhost:8080/api/classify-number")
}
