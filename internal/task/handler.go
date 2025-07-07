package task

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.POST("/tasks", h.CreateTask)
	e.GET("/tasks/:id", h.GetTask)
	e.GET("/tasks", h.GetAllTasks)
	e.PUT("/tasks/:id", h.UpdateTask)
	e.DELETE("/tasks/:id", h.DeleteTask)
}

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CreateTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	if err := h.Service.CreateTask(&task); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to create task"})
	}

	return c.JSON(201, task)
}

func (h *Handler) GetTask(c echo.Context) error {
	id := c.Param("id")
	task, err := h.Service.GetTaskByID(id)
	if err != nil {
		return c.NoContent(404)
	}
	return c.JSON(200, task)
}

func (h *Handler) GetAllTasks(c echo.Context) error {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to retrieve tasks"})
	}
	return c.JSON(200, tasks)
}

func (h *Handler) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	task.ID = id // Assuming Id is a string in Task struct
	if err := h.Service.UpdateTask(&task); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to update task"})
	}

	return c.JSON(200, task)
}

func (h *Handler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	if err := h.Service.DeleteTask(id); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to delete task"})
	}
	return c.NoContent(204)
}
