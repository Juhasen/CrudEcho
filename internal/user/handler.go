package user

import (
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
	var user User
	if err := c.Bind(&user); err != nil {
		return utils.ReturnApiError(c, http.StatusBadRequest, err)
	}

	if err := h.Service.CreateUser(&user); err != nil {
		return utils.ReturnApiError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
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
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	user.ID = id // Assuming ID is a string in User struct
	if err := h.Service.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if err := h.Service.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}
	return c.NoContent(http.StatusNoContent)
}
