package routes

import (
	"golang_API/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang Di WEB SERVICE ECHO")
	})
	e.GET("/produk", controllers.DataProduk)
	return e
}
