package router

import (
	"codesignal.com/todoapp/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()

	r.GET("/api/todos", controllers.GetAllTodos)

	r.GET("/api/todos/:id", controllers.GetTodoById)

	r.POST("/api/todos", controllers.AddTodoHandler)

	return r
}
