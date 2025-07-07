package main

import (
	"RestCrud/internal/task"
	"github.com/labstack/echo/v4"
	_ "net/http"

	"RestCrud/internal/user"
)

func main() {
	e := echo.New()

	// User domain setup
	userRepo := user.NewRepo()
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	user.RegisterRoutes(e, userHandler)

	// Task domain setup
	taskRepo := task.NewRepo()
	taskService := task.NewService(taskRepo)
	taskHandler := task.NewHandler(taskService)
	task.RegisterRoutes(e, taskHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
