package task

import (
	"RestCrud/pkg/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.POST("/api/tasks", h.CreateTask)
	e.GET("/api/tasks/:id", h.GetTask)
	e.GET("/api/tasks", h.GetAllTasks)
	e.PUT("/api/tasks/:id", h.UpdateTask)
	e.DELETE("/api/tasks/:id", h.DeleteTask)
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
		return utils.ReturnApiError(c, http.StatusBadRequest, err)
	}

	if err := h.Service.CreateTask(&task); err != nil {
		switch {
		case errors.Is(err, ErrAllArgumentsRequired):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrAllArgumentsRequired)
		case errors.Is(err, ErrInvalidStatus):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrInvalidStatus)
		case errors.Is(err, ErrInvalidDateFormat):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrInvalidDateFormat)
		case errors.Is(err, ErrDueDateInPast):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrDueDateInPast)
		case errors.Is(err, ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrTaskIdCannotBeEmpty)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusCreated, task)
}

func (h *Handler) GetTask(c echo.Context) error {
	id := c.Param("id")
	task, err := h.Service.GetTaskByID(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrTaskIdCannotBeEmpty)
		case errors.Is(err, ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, ErrTaskWithGivenIdNotFound)
		case errors.Is(err, ErrLoadDataFailed):
			return utils.ReturnApiError(c, http.StatusInternalServerError, ErrLoadDataFailed)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusOK, task)
}

func (h *Handler) GetAllTasks(c echo.Context) error {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		switch {
		case errors.Is(err, ErrLoadDataFailed):
			return utils.ReturnApiError(c, http.StatusInternalServerError, ErrLoadDataFailed)
		case errors.Is(err, ErrNoTasksFound):
			return utils.ReturnApiError(c, http.StatusNotFound, ErrNoTasksFound)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusOK, tasks)
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

	return c.JSON(http.StatusOK, task)
}

func (h *Handler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := h.Service.DeleteTask(id); err != nil {
		switch {
		case errors.Is(err, ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, 400, ErrTaskIdCannotBeEmpty)
		case errors.Is(err, ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, 404, ErrTaskWithGivenIdNotFound)
		}
		return c.JSON(500, map[string]string{"error": "Failed to delete task"})
	}
	return c.NoContent(http.StatusNoContent)
}
