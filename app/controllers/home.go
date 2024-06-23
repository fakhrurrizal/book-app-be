package controllers

import (
	"book-app/config"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.JSON(200, "Selamat Datang "+config.LoadConfig().AppName)
}