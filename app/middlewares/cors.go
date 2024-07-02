package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Cors Middleware
func Cors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowMethods: []string{echo.HEAD, echo.GET, echo.PUT, echo.POST, echo.DELETE},
	})
}
