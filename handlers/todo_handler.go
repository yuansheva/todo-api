package handlers

import (
	"net/http"
	"strconv"
	"todo-api/database"
	"todo-api/models"

	"github.com/labstack/echo/v4"
)

func GetAllTodo(c echo.Context) error {
	var todos []models.Todo
	database.DB.Find(&todos)

	return c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id")) // Konversi ID ke int
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	var todo models.Todo
	if err := database.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "To-Do not found"})
	}

	return c.JSON(http.StatusOK, todo)
}

func CreateTodo(c echo.Context) error {
	var todo models.Todo

	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request", "details": err.Error()})
	}

	if err := database.DB.Create(&todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create To-Do", "details": err.Error()})
	}

	return c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "To-Do not found"})
	}

	var updateData models.Todo
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Update hanya field yang ada
	database.DB.Model(&todo).Updates(updateData)

	return c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "To-Do not found"})
	}

	database.DB.Delete(&todo)

	return c.JSON(http.StatusOK, map[string]string{"message": "Todo has been deleted successfully"})
}
