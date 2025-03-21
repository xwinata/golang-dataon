package middlewares

import (
	"errors"
	"golang-dataon/customErrors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorSchema struct {
	Message string `json:"message"`
}

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			var apiErr *customErrors.ApiError
			if errors.As(err, &apiErr) {
				return c.JSON(apiErr.Code, ErrorSchema{
					Message: apiErr.Error(),
				})
			}

			return c.JSON(http.StatusInternalServerError, ErrorSchema{
				Message: err.Error(),
			})
		}
		return err
	}
}
