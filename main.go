package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos []Todo

// User is a simple struct representing a user.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Users is a slice of User representing a list of users.
var users = []User{
	{ID: 1, Name: "John"},
	{ID: 2, Name: "Jane"},
	{ID: 3, Name: "Doe"},
}

func main() {
	fmt.Printf("tes")
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Next()
	})

	// Define your API routes
	router.GET("/users", getUsers)
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", createTodo)
	router.DELETE("/todos/:id", deleteTodo)

	// Run the server on port 8080
	router.Run(":8080")
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
func getTodos(c *gin.Context) {

	c.JSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func createTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign a unique ID to the new todo
	newTodo.ID = len(todos) + 1

	// Append the new todo to the list
	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}
func deleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	fmt.Printf("Trying to delete todo with ID: %d\n", id)

	for i, todo := range todos {
		if todo.ID == id {
			// Delete the todo from the slice
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
