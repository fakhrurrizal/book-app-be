package router

import (
	"book-app/app/controllers"
	"book-app/app/middlewares"
	"log"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {

	app.Use(middlewares.Cors())
	app.Use(middlewares.Secure())

	app.GET("/", controllers.Index)

	app.Static("/assets", "assets")

	api := app.Group("/v1", middlewares.StripHTMLMiddleware)
	{
		book := api.Group("/book")
		{
			book.POST("", controllers.CreateBook)
			book.GET("/:id", controllers.GetBookByID)
			book.GET("", controllers.GetBooks)
			book.DELETE("/:id", controllers.DeleteBookByID)
			book.PUT("/:id", controllers.UpdateBookByID)
		}
		category := api.Group("/book-category")
		{
			category.POST("", controllers.CreateBookCategory)
			category.GET("/:id", controllers.GetBookCategoryByID)
			category.GET("", controllers.GetBookCategories)
			category.DELETE("/:id", controllers.DeleteBookCategoryByID)
			category.PUT("/:id", controllers.UpdateBookCategoryByID)
		}
		files := api.Group("/file")
		{
			files.POST("", controllers.UploadFile)
			files.GET("", controllers.GetFile)
		}
	}

	log.Printf("Server started...")
}
