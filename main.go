package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func test(c *gin.Context) {
	response := Response{
		Title:       "Test1",
		Description: "Hi, my name is Vincent Samuel Paul",
	}
	c.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := router.Group("/api/v1")
	{
		api.GET("/test", test)
	}

	log.Println("Go Server starting on port: 8000")
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
