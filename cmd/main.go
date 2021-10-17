package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	r := gin.Default()
	r.GET("/api/sample", getGinSample)
	r.GET("/api/todos", getTodos)

	r.Run(":" + port)
}

func getGinSample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H { "message": "Hello Gin World!" })
}

func getTodos(c *gin.Context) {
	type Todo struct {
		Id int
		Title string
	}

	todos := [...] Todo {
		{ Id: 1, Title: "TODO 1" },
		{ Id: 2, Title: "TODO 2" },
		{ Id: 3, Title: "TODO 3" },
	}

	c.JSON(http.StatusOK, gin.H { "todos": todos })
}
