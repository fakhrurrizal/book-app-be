package controllers

import (
	"book-app/app/repository"
	"book-app/app/reqres"
	"book-app/app/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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

func GetBooks(c echo.Context) error {
	categoryID, _ := strconv.Atoi(c.QueryParam("category_id"))
	param := utils.PopulatePaging(c, "status")

	data := repository.GetBooks(categoryID, param)

	return c.JSON(http.StatusOK, data)
}

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
