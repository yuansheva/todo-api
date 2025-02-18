package routes

import (
	"todo-api/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// todos application
	e.GET("/api/todos", handlers.GetAllTodo)
	e.GET("/api/todos/:id", handlers.GetTodoByID)
	e.POST("/api/todos", handlers.CreateTodo)
	e.PUT("/api/todos/:id", handlers.UpdateTodo)
	e.DELETE("/api/todos/:id", handlers.DeleteTodo)
}
