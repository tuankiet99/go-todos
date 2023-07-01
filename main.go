package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Learn Golang", Completed: false},
	{ID: "2", Item: "Workout", Completed: true},
	{ID: "3", Item: "Read book", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusOK, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById((id))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func updateStatusTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func removeTodo(context *gin.Context) {
	id := context.Param("id")

	var indexRemove int
	for index, todo := range todos {
		if todo.ID == id {
			fmt.Println("todoID, id", todo.ID, id)
			indexRemove = index

			result := append(todos[:indexRemove], todos[indexRemove+1:]...)
			todos = result
			context.IndentedJSON(http.StatusOK, todos)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}
func getTodoById(id string) (*todo, error) {
	for index, todo := range todos {
		if todo.ID == id {
			return &todos[index], nil
		}
	}
	return nil, errors.New("Todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.PATCH("/todos/:id", updateStatusTodo)
	router.DELETE("/todos/:id", removeTodo)
	router.Run("localhost:9000")
}
