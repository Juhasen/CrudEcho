package utils

import "github.com/labstack/echo/v4"

func ReturnApiError(c echo.Context, status int, err error) error {
	return c.JSON(status, map[string]string{"error": err.Error()})
}
