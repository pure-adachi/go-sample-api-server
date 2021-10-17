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

	r.Run(":" + port)
}

func getGinSample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H { "message": "Hello Gin World!" })
}