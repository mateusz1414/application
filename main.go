package main

import (
	"application/pages"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.Static("/assets", "./assets/css")
	server.GET("/:p", pages.IndexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)

}
