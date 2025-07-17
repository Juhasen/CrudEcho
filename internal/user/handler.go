package user

import (
	generated "RestCrud/openapi"
	"RestCrud/pkg/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	Service *Service
}

func NewHandler(service *Service) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var userRequest generated.UserResponse

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

func (h *UserHandler) GetUserById(c echo.Context) error {
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

func (h *UserHandler) GetAllUsers(c echo.Context) error {
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

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ReturnApiError(c, http.StatusBadRequest, ErrUserIDRequired)
	}

	var user generated.UserRequest

	if err := c.Bind(&user); err != nil {
		return utils.ReturnApiError(c, http.StatusBadRequest, err)
	}

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

func (h *UserHandler) DeleteUser(c echo.Context) error {
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
