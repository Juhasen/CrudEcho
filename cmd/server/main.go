package main

import (
	"RestCrud/internal/task"
	taskModel "RestCrud/internal/task/model"
	"RestCrud/internal/user"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	_ "net/http"

	userModel "RestCrud/internal/user/model"
)

func main() {
	e := echo.New()

	// Initialize DB connection
	dsn := "host=localhost user=postgres password=123 dbname=restcrud port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Connected to database")

	if err := db.AutoMigrate(&userModel.User{}, &taskModel.Task{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migration completed successfully")

	userRepo := user.NewRepo(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	user.RegisterRoutes(e, userHandler)

	log.Println("User domain setup completed")

	taskRepo := task.NewRepo(db)
	taskService := task.NewService(taskRepo, userRepo)
	taskHandler := task.NewHandler(taskService)
	task.RegisterRoutes(e, taskHandler)

	log.Println("Task domain setup completed")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	log.Println("Server started on port 1323")
}
