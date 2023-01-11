package controllers

import (
	"golang_API/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SigninUser(c echo.Context) error {
	result, err := models.SigninUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
