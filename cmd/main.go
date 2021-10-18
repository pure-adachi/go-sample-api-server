package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	r := gin.Default()
	r.GET("/api/sample", getGinSample)
	r.GET("/api/todos", getTodos)
	r.GET("/api/todos/:Id", getTodo)
	r.POST("/api/todos", addTodo)
	r.PATCH("/api/todos/:Id", updateTodo)
	r.DELETE("/api/todos/:Id", deleteTodo)

	r.Run(":" + port)
}

func getGinSample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H { "message": "Hello Gin World!" })
}

type Todo struct {
	Id int
	Title string
}

func todos() [3]Todo {
	return [3] Todo {
		{ Id: 1, Title: "TODO 1" },
		{ Id: 2, Title: "TODO 2" },
		{ Id: 3, Title: "TODO 3" },
	}
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H { "todos": todos() })
}

func getTodo(c *gin.Context) {
	Id := c.Param("Id")
	id, _ := strconv.Atoi(Id)

	todo := findTodoById(id)

	c.JSON(http.StatusOK, gin.H { "todo": todo })
}

type InputTodo struct {
	Title string
}

func addTodo(c *gin.Context) {
	var inputTodo InputTodo
	c.BindJSON(&inputTodo)

	data := todos()
	count := len(data)
	newTodo := Todo { Id: count + 1, Title: inputTodo.Title }

	c.JSON(http.StatusOK, gin.H { "todo": newTodo })
}

func updateTodo(c *gin.Context) {
	Id := c.Param("Id")
	id, _ := strconv.Atoi(Id)

	todo := findTodoById(id)

	var inputTodo InputTodo
	c.BindJSON(&inputTodo)

	if todo.Id != -1 {
		todo.Title = inputTodo.Title
	}

	c.JSON(http.StatusOK, gin.H { "todo": todo })
}

func deleteTodo(c *gin.Context) {
	Id := c.Param("Id")
	id, _ := strconv.Atoi(Id)

	var activeTodo [] Todo

	for _, todo := range todos() {
		if todo.Id != id {
			activeTodo = append(activeTodo, todo)
		}
	}
	
	c.JSON(http.StatusOK, gin.H { "todos": activeTodo })
}

func findTodoById(id int) Todo {
	for _, todo := range todos() {
		if todo.Id == id {
			return todo
		}
	}

	return Todo{ Id: -1 }
}
