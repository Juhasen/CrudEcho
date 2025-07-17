package main

import (
	"RestCrud/internal"
	"RestCrud/internal/model"
	"RestCrud/internal/task"
	"RestCrud/internal/user"
	"RestCrud/openapi"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	_ "net/http"
	"os"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	// Get connection string from environment variable
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN environment variable is not set")
	}
	// Initialize DB connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Connected to database")

	if err := db.AutoMigrate(&model.User{}, &model.Task{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migration completed successfully")

	userRepo := user.NewRepo(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	log.Println("User domain setup completed")

	taskRepo := task.NewRepo(db)
	taskService := task.NewService(taskRepo, userRepo)
	taskHandler := task.NewHandler(taskService)

	log.Println("Task domain setup completed")

	apiGroup := e.Group("/api")

	handler := internal.NewHandler(userHandler, taskHandler)

	generated.RegisterHandlers(apiGroup, generated.ServerInterface(handler))

	e.Static("/swagger-ui", "swagger-ui")
	e.File("/openapi.yml", "dist/openapi.yml")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	log.Println("Server started on port 1323")
}
