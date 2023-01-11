package controllers

import (
	"golang_API/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DataKategori(c echo.Context) error {
	result, err := models.DataKategori()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
