package router

import (
	"book-app/app/controllers"
	"book-app/app/middlewares"
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	// renderer := &TemplateRenderer{
	// 	templates: template.Must(template.ParseGlob("*.html")),
	// }
	// app.Renderer = renderer
	app.Use(middlewares.Cors())
	app.Use(middlewares.Secure())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Recover())

	// app.GET("/swagger/*", echoSwagger.WrapHandler)

	// app.GET("/docs", func(c echo.Context) error {
	// 	err := c.Render(http.StatusOK, "docs.html", map[string]interface{}{
	// 		"BaseUrl": config.LoadConfig().BaseUrl,
	// 		"Title":   "Dokumentasi API " + config.LoadConfig().AppName,
	// 	})
	// 	if err != nil {
	// 		fmt.Println("Error rendering docs.html:", err)
	// 	}
	// 	return err
	// })

	app.GET("/", controllers.Index)

	app.Static("/assets", "assets")

	api := app.Group("/v1", middlewares.StripHTMLMiddleware)
	{
		auth := api.Group("/auth")
		{
			auth.POST("/signin", controllers.SignIn)
			auth.POST("/signup", controllers.SignUp)
			auth.GET("/user", controllers.GetSignInUser, middlewares.Auth())
		}
		book := api.Group("/book")
		{
			book.POST("", controllers.UploadFile(controllers.CreateBook), middlewares.Auth())
			book.GET("/:id", controllers.GetBookByID)
			book.GET("", controllers.GetBooks)
			book.DELETE("/:id", controllers.DeleteBookByID, middlewares.Auth())
			book.PUT("/:id", controllers.UploadFile(controllers.UpdateBookByID), middlewares.Auth())
		}
		category := api.Group("/book-category")
		{
			category.POST("", controllers.CreateBookCategory, middlewares.Auth())
			category.GET("/:id", controllers.GetBookCategoryByID)
			category.GET("", controllers.GetBookCategories)
			category.DELETE("/:id", controllers.DeleteBookCategoryByID, middlewares.Auth())
			category.PUT("/:id", controllers.UpdateBookCategoryByID, middlewares.Auth())
		}
		book_lending := api.Group("/book-lending")
		{
			book_lending.POST("", controllers.CreateBookLending, middlewares.Auth())
			book_lending.GET("/:id", controllers.GetBookLendingByID)
			book_lending.GET("", controllers.GetBookLendings)
			book_lending.DELETE("/:id", controllers.DeleteBookLendingByID, middlewares.Auth())
			book_lending.PUT("/:id", controllers.UpdateBookLendingByID, middlewares.Auth())
		}
	}

	log.Printf("Server started...")
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
