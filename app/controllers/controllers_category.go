package controllers

import (
	"book-app/app/repository"
	"book-app/app/reqres"
	"book-app/app/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateBookCategory(c echo.Context) error {
	var input reqres.BookCategoryRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}

	utils.StripTagsFromStruct(&input)

	data, err := repository.CreateBookCategory(&input)

	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Gagal Menambahkan Data"))
	}

	response, err := repository.GetBookCategoryID(int(data.ID))
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Gagal Mendapatkan Data"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    response,
		"message": "Tambah Kategori Berhasil",
	})
}

func GetBookCategories(c echo.Context) error {

	param := utils.PopulatePaging(c, "status")
	data := repository.GetBookCategories(param)

	return c.JSON(http.StatusOK, data)
}

func GetBookCategoryByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repository.GetBookCategoryID(id)
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Gagal mendapatkan data"))
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "Berhasil mendapatkan data",
	})
}

func DeleteBookCategoryByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	check, _ := repository.GetAllBooksPlain("", id)

	if len(check) > 0 {
		return c.JSON(400, utils.Respond(400, "bad request", "Kategori ini tidak dapat dihapus karena sudah digunakan"))
	}

	data, err := repository.GetBookCategoryIDPlain(id)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Gagal mendapatkan data"))
	}

	response, err := repository.GetBookCategoryID(int(data.ID))
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Gagal mendapatkan response"))
	}

	_, err = repository.DeleteBookCategory(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Gagal menghapus data"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    response,
		"message": "Berhasil menghapus data",
	})
}

func UpdateBookCategoryByID(c echo.Context) error {
	var input reqres.BookCategoryRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}
	utils.StripTagsFromStruct(&input)

	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repository.GetBookCategoryIDPlain(id)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Gagal mnedapatkan data"))
	}

	if input.Name != "" {
		data.Name = input.Name
	}
	if input.Description != "" {
		data.Description = input.Description
	}
	data.Status = input.Status

	dataUpdate, err := repository.UpdateBookCategory(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Gagal menghapus data"))
	}

	response, err := repository.GetBookCategoryID(int(dataUpdate.ID))
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Gagal mendapatkan data response"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    response,
		"message": "Berhasil Mengubah data",
	})
}
