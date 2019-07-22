package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()

	r.GET("/todo", addTodoItem)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}

func addTodoItem(c *gin.Context) {

	c.JSON(200, "abc")
}
