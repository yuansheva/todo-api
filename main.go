package main

import (
	"todo-api/database"
	"todo-api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "todo-api/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize database
	database.InitDB()

	// setup swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Setup routes
	routes.SetupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
