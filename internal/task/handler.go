package task

import (
	"RestCrud/internal/task/dto"
	taskErr "RestCrud/internal/task/errors"
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
		case errors.Is(err, taskErr.ErrAllArgumentsRequired):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrAllArgumentsRequired)
		case errors.Is(err, taskErr.ErrInvalidStatus):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrInvalidStatus)
		case errors.Is(err, taskErr.ErrInvalidDateFormat):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrInvalidDateFormat)
		case errors.Is(err, taskErr.ErrDueDateInPast):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrDueDateInPast)
		case errors.Is(err, taskErr.ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrTaskIdCannotBeEmpty)
		case errors.Is(err, taskErr.ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, taskErr.ErrTaskWithGivenIdNotFound)
		case errors.Is(err, taskErr.ErrUserWithGivenIdDoesNotExist):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrUserWithGivenIdDoesNotExist)
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
		case errors.Is(err, taskErr.ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrTaskIdCannotBeEmpty)
		case errors.Is(err, taskErr.ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, taskErr.ErrTaskWithGivenIdNotFound)
		case errors.Is(err, taskErr.ErrLoadDataFailed):
			return utils.ReturnApiError(c, http.StatusInternalServerError, taskErr.ErrLoadDataFailed)
		case errors.Is(err, taskErr.ErrIdIsNotValid):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrIdIsNotValid)
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
		case errors.Is(err, taskErr.ErrLoadDataFailed):
			return utils.ReturnApiError(c, http.StatusInternalServerError, taskErr.ErrLoadDataFailed)
		case errors.Is(err, taskErr.ErrNoTasksFound):
			return utils.ReturnApiError(c, http.StatusNotFound, taskErr.ErrNoTasksFound)
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
		case errors.Is(err, taskErr.ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrTaskIdCannotBeEmpty)
		case errors.Is(err, taskErr.ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, taskErr.ErrTaskWithGivenIdNotFound)
		case errors.Is(err, taskErr.ErrInvalidStatus):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrInvalidStatus)
		case errors.Is(err, taskErr.ErrInvalidDateFormat):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrInvalidDateFormat)
		case errors.Is(err, taskErr.ErrDueDateInPast):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrDueDateInPast)
		case errors.Is(err, taskErr.ErrIdIsNotValid):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrIdIsNotValid)
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
		case errors.Is(err, taskErr.ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrTaskIdCannotBeEmpty)
		case errors.Is(err, taskErr.ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, taskErr.ErrTaskWithGivenIdNotFound)
		case errors.Is(err, taskErr.ErrIdIsNotValid):
			return utils.ReturnApiError(c, http.StatusBadRequest, taskErr.ErrIdIsNotValid)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}
	return c.NoContent(http.StatusNoContent)
}
