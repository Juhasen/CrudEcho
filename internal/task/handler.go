package task

import (
	generated "RestCrud/openapi"
	"RestCrud/pkg/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskHandler struct {
	Service *Service
}

func NewHandler(service *Service) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var task generated.TaskRequest
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
		case errors.Is(err, ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, ErrTaskWithGivenIdNotFound)
		case errors.Is(err, ErrUserWithGivenIdDoesNotExist):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrUserWithGivenIdDoesNotExist)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTaskById(c echo.Context) error {
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
		case errors.Is(err, ErrIdIsNotValid):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrIdIsNotValid)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) GetAllTasks(c echo.Context) error {
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

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	var task generated.TaskRequest
	if err := c.Bind(&task); err != nil {
		return utils.ReturnApiError(c, http.StatusBadRequest, err)
	}

	if err := h.Service.UpdateTask(id, &task); err != nil {
		switch {
		case errors.Is(err, ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrTaskIdCannotBeEmpty)
		case errors.Is(err, ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, ErrTaskWithGivenIdNotFound)
		case errors.Is(err, ErrInvalidStatus):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrInvalidStatus)
		case errors.Is(err, ErrInvalidDateFormat):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrInvalidDateFormat)
		case errors.Is(err, ErrDueDateInPast):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrDueDateInPast)
		case errors.Is(err, ErrIdIsNotValid):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrIdIsNotValid)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c echo.Context, id string) error {
	if err := h.Service.DeleteTask(id); err != nil {
		switch {
		case errors.Is(err, ErrTaskIdCannotBeEmpty):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrTaskIdCannotBeEmpty)
		case errors.Is(err, ErrTaskWithGivenIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, ErrTaskWithGivenIdNotFound)
		case errors.Is(err, ErrIdIsNotValid):
			return utils.ReturnApiError(c, http.StatusBadRequest, ErrIdIsNotValid)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}
	return c.NoContent(http.StatusNoContent)
}
