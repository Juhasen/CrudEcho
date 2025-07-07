package user

import "github.com/labstack/echo/v4"

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
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	if err := h.Service.CreateUser(&user); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(201, user)
}

func (h *Handler) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.Service.GetUserByID(id)
	if err != nil {
		return c.NoContent(404)
	}
	return c.JSON(200, user)
}

func (h *Handler) GetAllUsers(c echo.Context) error {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to retrieve users"})
	}
	return c.JSON(200, users)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	user.ID = id // Assuming ID is a string in User struct
	if err := h.Service.UpdateUser(&user); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to update user"})
	}

	return c.JSON(200, user)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if err := h.Service.DeleteUser(id); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to delete user"})
	}
	return c.NoContent(204)
}
