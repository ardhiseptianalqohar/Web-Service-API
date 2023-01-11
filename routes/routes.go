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

	//UNTUK MENAMPILKAN ATAU MENGAMBIL DATA DARI DATABASE
	e.GET("/produk", controllers.DataProduk)

	e.GET("/signin", controllers.SigninUser)

	e.GET("/kategori", controllers.DataKategori)

	return e
}
