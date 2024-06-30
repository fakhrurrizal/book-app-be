package controllers

import (
	"book-app/app/repository"
	"book-app/app/reqres"
	"book-app/app/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateBook godoc
// @Summary Create Book
// @Description Create New Book
// @Tags Book
// @Produce json
// @Param Body body reqres.BookRequest true "Create body"
// @Success 200
// @Router /v1/book [post]
func CreateBook(c echo.Context) error {
	var input reqres.BookRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}

	utils.StripTagsFromStruct(&input)

	data, err := repository.CreateBook(&input)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to create"))
	}

	response, err := repository.GetBookByID(int(data.ID))
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Failed to get response"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    response,
		"message": "Berhasil menambah data buku",
	})
}

// GetBookData godoc
// @Summary Get Book with Pagination
// @Description Get Book with Pagination
// @Tags Book
// @Param search query string false "search (string)"
// @Param page query integer false "page (int)"
// @Param limit query integer false "limit (int)"
// @Param sort query integer false "sort (id or publication_year or title)"
// @Param order query integer false "order (asc or desc)"
// @Param status query integer false "status (status)"
// @Param category_id query integer false "category_id (int)"
// @Produce json
// @Success 200
// @Router /v1/book [get]
func GetBooks(c echo.Context) error {
	categoryID, _ := strconv.Atoi(c.QueryParam("category_id"))
	param := utils.PopulatePaging(c, "status")

	data := repository.GetBooks(categoryID, param)

	return c.JSON(http.StatusOK, data)
}

// GetBookId godoc
// @Summary Get Single Book
// @Description Get Single Book
// @Tags Book
// @Param id path string true "id"
// @Produce json
// @Success 200
// @Router /v1/book/{id} [get]
func GetBookByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := repository.GetBookByID(id)
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Data tidak tersedia"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "Berhasil mendapatkan data",
	})
}

// DeleteBookId godoc
// @Summary Delete Single Book by ID
// @Description Delete Single Book by ID
// @Tags Book
// @Produce json
// @Param id path integer true "id"
// @Success 200
// @Router /v1/book/{id} [delete]
func DeleteBookByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := repository.GetBookByIDPlain(id)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Data tidak tersedia"))
	}

	_, err = repository.DeleteBook(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Gagal menghapus data"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "Berhasil menghapus data",
	})

}

// UpdateBookById godoc
// @Summary Update Single Book by ID
// @Description Update Single Book by ID
// @Tags Book
// @Produce json
// @Param id path integer true "id"
// @Param Body body reqres.BookRequest true "Update body"
// @Success 200
// @Router /v1/book/{id} [put]
func UpdateBookByID(c echo.Context) error {
	var input reqres.BookRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}

	utils.StripTagsFromStruct(&input)

	id, _ := strconv.Atoi(c.Param("id"))

	data, err := repository.GetBookByIDPlain(id)

	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Data tidak tersedia"))
	}

	if input.Title != "" {
		data.Title = input.Title
	}

	if input.Author != "" {
		data.Author = input.Author
	}

	if input.Description != "" {
		data.Description = input.Description
	}

	if input.Image != "" {
		data.Image = input.Image
	}

	if input.Publisher != "" {
		data.Publisher = input.Publisher
	}

	if input.BookCode != "" {
		data.BookCode = input.BookCode
	}

	if input.Language != "" {
		data.Language = input.Language
	}

	if input.CategoryID != 0 {
		data.CategoryID = input.CategoryID
	}

	if input.NumberOfPages != 0 {
		data.NumberOfPages = input.NumberOfPages
	}

	if input.PublicationYear != 0 {
		data.PublicationYear = input.PublicationYear
	}

	data.Status = input.Status

	dataUpdate, err := repository.UpdateBook(data)

	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Gagal mengubah data"))
	}

	response, err := repository.GetBookByID(int(dataUpdate.ID))
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Failed to get response"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    response,
		"message": "Success to update",
	})
}
