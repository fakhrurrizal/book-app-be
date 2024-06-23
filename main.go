package main

import (
	router "book-app/app/routers"
	"book-app/config"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/tylerb/graceful.v1"
)

func main() {
	app := echo.New()

	config.Database()

	router.Init(app)

	app.Server.Addr = "127.0.0.1:" + config.LoadConfig().Port
	log.Printf("Server: " + config.LoadConfig().BaseUrl)

	graceful.ListenAndServe(app.Server, 5*time.Second)
}
