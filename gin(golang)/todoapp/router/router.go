package router

import (
	"codesignal.com/todoapp/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()

	r.GET("/api/todos", controllers.GetAllTodos)

	r.GET("/api/todos/:id", controllers.GetTodoById)

	return r
}
