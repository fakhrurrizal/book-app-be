package router

import (
	"book-app/app/controllers"
	"book-app/app/middlewares"
	_ "book-app/docs"
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
		book := api.Group("/book")
		{
			book.POST("", middlewares.UploadFile(controllers.CreateBook))
			book.GET("/:id", controllers.GetBookByID)
			book.GET("", controllers.GetBooks)
			book.DELETE("/:id", controllers.DeleteBookByID)
			book.PUT("/:id", middlewares.UploadFile(controllers.UpdateBookByID))
		}
		category := api.Group("/book-category")
		{
			category.POST("", controllers.CreateBookCategory)
			category.GET("/:id", controllers.GetBookCategoryByID)
			category.GET("", controllers.GetBookCategories)
			category.DELETE("/:id", controllers.DeleteBookCategoryByID)
			category.PUT("/:id", controllers.UpdateBookCategoryByID)
		}
		// files := api.Group("/file")
		// {
		// 	files.POST("", controllers.UploadFile)
		// 	files.GET("", controllers.GetFile)
		// }
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
