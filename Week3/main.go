package main

import (
	"example/web-service-gin/todos"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello")
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", postTodos)
	router.GET("/todos/:id", getTodoByID)
	router.PATCH("/todos/:id", updateTodoByID)
	router.DELETE("/todos/:id", deleteTodoByID)
	router.Run("localhost:3002")
}

var todosList = []todos.Todo{
	{UserId: 1, Id: 1, Title: "delectus aut autem", Completed: false},
	{UserId: 1, Id: 2, Title: "quis ut nam facilis et officia qui", Completed: false},
	{UserId: 1, Id: 3, Title: "fugiat veniam minus", Completed: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todosList)
}

func postTodos(c *gin.Context) {
	var newTodo todos.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todosList = append(todosList, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range todosList {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "TODO not found"})
}

func updateTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.

	newTodo := todos.Todo{}

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		// Nếu không thể bind JSON, trả về lỗi
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(newTodo)

	for index, a := range todosList {
		if a.Id == id {
			todosList[index].Title = newTodo.Title
			todosList[index].Completed = newTodo.Completed
			c.IndentedJSON(http.StatusOK, todosList[index])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "TODO not found"})
}

func deleteTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.

	for index, a := range todosList {
		if a.Id == id {
			todosList = append(todosList[:index], todosList[index+1:]...)
			c.IndentedJSON(http.StatusOK, todosList)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "TODO not found"})
}
