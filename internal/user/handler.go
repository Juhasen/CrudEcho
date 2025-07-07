package user

import (
	"RestCrud/internal/user/dto"
	"RestCrud/pkg/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.POST("/users", h.CreateUser)
	e.GET("/users/:id", h.GetUser)
	e.GET("/users", h.GetAllUsers)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)
}

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CreateUser(c echo.Context) error {
	var userRequest dto.UserDTO

	if err := c.Bind(&userRequest); err != nil {
		return utils.ReturnApiError(c, http.StatusBadRequest, err)
	}

	if err := h.Service.CreateUser(&userRequest); err != nil {
		return utils.ReturnApiError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, userRequest)
}

func (h *Handler) GetUser(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ReturnApiError(c, http.StatusBadRequest, ErrUserIDRequired)
	}

	user, err := h.Service.GetUserByID(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrUserIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, err)
		case errors.Is(err, ErrUserIDRequired):
			return utils.ReturnApiError(c, http.StatusBadRequest, err)
		case errors.Is(err, ErrLoadDataFailed):
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}

		return utils.ReturnApiError(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) GetAllUsers(c echo.Context) error {
	users, err := h.Service.GetAllUsers()

	if err != nil {
		switch {
		case errors.Is(err, ErrNoUsersFound):
			return utils.ReturnApiError(c, http.StatusNotFound, err)
		case errors.Is(err, ErrUserIDRequired):
			return utils.ReturnApiError(c, http.StatusBadRequest, err)
		case errors.Is(err, ErrLoadDataFailed):
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}

		return utils.ReturnApiError(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ReturnApiError(c, http.StatusBadRequest, ErrUserIDRequired)
	}

	var user dto.UserUpdateDTO

	if err := c.Bind(&user); err != nil {
		return utils.ReturnApiError(c, http.StatusBadRequest, err)
	}

	user.ID = id

	if err := h.Service.UpdateUser(id, &user); err != nil {
		return utils.ReturnApiError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ReturnApiError(c, http.StatusBadRequest, ErrUserIDRequired)
	}

	if err := h.Service.DeleteUser(id); err != nil {
		return utils.ReturnApiError(c, http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
