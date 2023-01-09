package controllers

import (
	"golang_API/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DataProduk(c echo.Context) error {
	result, err := models.DataProduk()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
