package handler

import (
	router "book-app/app/routers"
	"book-app/config"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/tylerb/graceful.v1"
)

func main() {
	app := Start()

	app.Server.Addr = "127.0.0.1:" + config.LoadConfig().Port
	log.Printf("Server: " + config.LoadConfig().BaseUrl)

	graceful.ListenAndServe(app.Server, 5*time.Second)
}

func Main(w http.ResponseWriter, r *http.Request) {
	e := Start()

	e.ServeHTTP(w, r)
}

func Start() *echo.Echo {
	app := echo.New()

	config.Database()

	router.Init(app)

	return app
}
