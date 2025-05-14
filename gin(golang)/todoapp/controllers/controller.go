package controllers

import (
	"net/http"
	"strconv"

	"codesignal.com/todoapp/services"
	"github.com/gin-gonic/gin"
)

func GetAllTodos(c *gin.Context) {
	completedQuery := c.Query("completed")
	var completed *bool = nil

	if completedQuery != "" {
		value, err := strconv.ParseBool(completedQuery)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'completed' query parameter. Use true or false."})
		}
		completed = &value
	}

	todo := services.FilterTodo(completed)

	c.JSON(http.StatusOK, todo)
}

func GetTodoById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	todo, fnd := services.FilterTodoById(id)

	if fnd {
		c.JSON(http.StatusOK, todo)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	}

}
