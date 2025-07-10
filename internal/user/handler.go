package user

import (
	"RestCrud/pkg/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.POST("/api/users", h.CreateUser)
	e.GET("/api/users/:id", h.GetUser)
	e.GET("/api/users", h.GetAllUsers)
	e.PUT("/api/users/:id", h.UpdateUser)
	e.DELETE("/api/users/:id", h.DeleteUser)
}

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CreateUser(c echo.Context) error {
	var userRequest ResponseDTO

	if err := c.Bind(&userRequest); err != nil {
		return utils.ReturnApiError(c, http.StatusBadRequest, err)
	}

	if err := h.Service.CreateUser(&userRequest); err != nil {
		switch {
		case errors.Is(err, ErrUserNameRequired):
			return utils.ReturnApiError(c, http.StatusBadRequest, err)
		case errors.Is(err, ErrUserEmailRequired):
			return utils.ReturnApiError(c, http.StatusBadRequest, err)
		case errors.Is(err, ErrUserEmailInvalid):
			return utils.ReturnApiError(c, http.StatusBadRequest, err)
		case errors.Is(err, ErrUserAlreadyExists):
			return utils.ReturnApiError(c, http.StatusConflict, err)
		case errors.Is(err, ErrUserIDRequired):
			return utils.ReturnApiError(c, http.StatusBadRequest, err)
		case errors.Is(err, ErrLoadDataFailed):
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
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
		case errors.Is(err, ErrIdIsNotValid):
			return c.NoContent(http.StatusBadRequest)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
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

	var user RequestDTO

	if err := c.Bind(&user); err != nil {
		return utils.ReturnApiError(c, http.StatusBadRequest, err)
	}

	user.ID = id

	if err := h.Service.UpdateUser(id, &user); err != nil {
		switch {
		case errors.Is(err, ErrUserIDRequired):
			return utils.ReturnApiError(c, http.StatusBadRequest, err)
		case errors.Is(err, ErrUserIDMismatch):
			return utils.ReturnApiError(c, http.StatusBadRequest, err)
		case errors.Is(err, ErrUserAlreadyExists):
			return utils.ReturnApiError(c, http.StatusConflict, err)
		case errors.Is(err, ErrAtLeastOneFieldRequired):
			return utils.ReturnApiError(c, http.StatusBadRequest, err)
		case errors.Is(err, ErrIdIsNotValid):
			return c.NoContent(http.StatusBadRequest)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ReturnApiError(c, http.StatusBadRequest, ErrUserIDRequired)
	}

	if err := h.Service.DeleteUser(id); err != nil {
		switch {
		case errors.Is(err, ErrUserIDRequired):
			return utils.ReturnApiError(c, http.StatusBadRequest, err)
		case errors.Is(err, ErrUserIdNotFound):
			return utils.ReturnApiError(c, http.StatusNotFound, err)
		case errors.Is(err, ErrFailedToDeleteUser):
			return utils.ReturnApiError(c, http.StatusConflict, err)
		case errors.Is(err, ErrDeleteUserNotFound):
			return utils.ReturnApiError(c, http.StatusBadRequest, err)
		case errors.Is(err, ErrIdIsNotValid):
			return c.NoContent(http.StatusBadRequest)
		default:
			return utils.ReturnApiError(c, http.StatusInternalServerError, err)
		}
	}

	return c.NoContent(http.StatusNoContent)
}
