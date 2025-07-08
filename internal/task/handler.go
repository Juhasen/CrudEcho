package task

import (
	"RestCrud/internal/task/dto"
	errors2 "RestCrud/internal/task/errors"
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
	var task dto.TaskRequestDTO
	if err := c.Bind(&task); err != nil {
		return utils.ReturnApiError(c, http.StatusBadRequest, err)
	}

	if err := h.Service.CreateTask(&task); err != nil {
		switch {
		case errors.Is(err, errors2.ErrAllArgumentsRequired):
			return utils.ReturnApiError(c, http.StatusBadRequest, errors2.ErrAllArgumentsRequired)
		case errors.Is(err, errors2.ErrInvalidStatus):
			return utils.ReturnApiError(c, http.StatusBadRequest, errors2.ErrInvalidStatus)
		case errors.Is(err, errors2.ErrInvalidDateFormat):
			return utils.ReturnApiError(c, http.StatusBadRequest, errors2.ErrInvalidDateFormat)
		case errors.Is(err, errors2.ErrDueDateInPast):
			return utils.ReturnApiError(c, http.StatusBadRequest, errors2.ErrDueDateInPast)
		case errors.Is(err, errors2.ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, http.StatusBadRequest, errors2.ErrTaskIdCannotBeEmpty)
		case errors.Is(err, errors2.ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, errors2.ErrTaskWithGivenIdNotFound)
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
		case errors.Is(err, errors2.ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, http.StatusBadRequest, errors2.ErrTaskIdCannotBeEmpty)
		case errors.Is(err, errors2.ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, errors2.ErrTaskWithGivenIdNotFound)
		case errors.Is(err, errors2.ErrLoadDataFailed):
			return utils.ReturnApiError(c, http.StatusInternalServerError, errors2.ErrLoadDataFailed)
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
		case errors.Is(err, errors2.ErrLoadDataFailed):
			return utils.ReturnApiError(c, http.StatusInternalServerError, errors2.ErrLoadDataFailed)
		case errors.Is(err, errors2.ErrNoTasksFound):
			return utils.ReturnApiError(c, http.StatusNotFound, errors2.ErrNoTasksFound)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *Handler) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	var task dto.TaskRequestDTO
	if err := c.Bind(&task); err != nil {
		return utils.ReturnApiError(c, http.StatusBadRequest, err)
	}

	if err := h.Service.UpdateTask(id, &task); err != nil {
		switch {
		case errors.Is(err, errors2.ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, http.StatusBadRequest, errors2.ErrTaskIdCannotBeEmpty)
		case errors.Is(err, errors2.ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, errors2.ErrTaskWithGivenIdNotFound)
		case errors.Is(err, errors2.ErrInvalidStatus):
			return utils.ReturnApiError(c, http.StatusBadRequest, errors2.ErrInvalidStatus)
		case errors.Is(err, errors2.ErrInvalidDateFormat):
			return utils.ReturnApiError(c, http.StatusBadRequest, errors2.ErrInvalidDateFormat)
		case errors.Is(err, errors2.ErrDueDateInPast):
			return utils.ReturnApiError(c, http.StatusBadRequest, errors2.ErrDueDateInPast)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}

	return c.JSON(http.StatusOK, task)
}

func (h *Handler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := h.Service.DeleteTask(id); err != nil {
		switch {
		case errors.Is(err, errors2.ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, http.StatusBadRequest, errors2.ErrTaskIdCannotBeEmpty)
		case errors.Is(err, errors2.ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, errors2.ErrTaskWithGivenIdNotFound)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}
	return c.NoContent(http.StatusNoContent)
}
