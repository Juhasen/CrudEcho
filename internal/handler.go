package internal

import (
	"RestCrud/internal/task"
	"RestCrud/internal/user"
	"github.com/labstack/echo/v4"
	openapitypes "github.com/oapi-codegen/runtime/types"
)

type ServerInterface interface {
	GetTasks(c echo.Context) error
	CreateTask(c echo.Context) error
	DeleteTask(c echo.Context, taskId openapitypes.UUID) error
	GetTaskById(c echo.Context, taskId openapitypes.UUID) error
	UpdateTask(c echo.Context, taskId openapitypes.UUID) error

	GetUsers(c echo.Context) error
	CreateUser(c echo.Context) error
	DeleteUser(c echo.Context, userId openapitypes.UUID) error
	GetUserById(c echo.Context, userId openapitypes.UUID) error
	UpdateUser(c echo.Context, userId openapitypes.UUID) error
}

type Handler struct {
	userHandler *user.UserHandler
	taskHandler *task.TaskHandler
}

func NewHandler(u *user.UserHandler, t *task.TaskHandler) ServerInterface {
	return &Handler{
		userHandler: u,
		taskHandler: t,
	}
}

// Metody task'u

func (h *Handler) GetTasks(c echo.Context) error {
	return h.taskHandler.GetAllTasks(c)
}

func (h *Handler) CreateTask(c echo.Context) error {
	return h.taskHandler.CreateTask(c)
}

func (h *Handler) DeleteTask(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.taskHandler.DeleteTask(c, id.String())
}

func (h *Handler) GetTaskById(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.taskHandler.GetTaskById(c)
}

func (h *Handler) UpdateTask(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.taskHandler.UpdateTask(c)
}

// Metody user'a

func (h *Handler) GetUsers(c echo.Context) error {
	return h.userHandler.GetAllUsers(c)
}

func (h *Handler) CreateUser(c echo.Context) error {
	return h.userHandler.CreateUser(c)
}

func (h *Handler) DeleteUser(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.userHandler.DeleteUser(c)
}

func (h *Handler) GetUserById(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.userHandler.GetUserById(c)
}

func (h *Handler) UpdateUser(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.userHandler.UpdateUser(c)
}
