package controllers

import (
	"book-app/app/repository"
	"book-app/app/reqres"
	"book-app/app/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateBookCategory godoc
// @Summary Create Book Category
// @Description Create New Book Category
// @Tags BookCategory
// @Produce json
// @Param Body body reqres.BookCategoryRequest true "Create body"
// @Success 200
// @Router /v1/book-category [post]
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

// GetCategoryBookData godoc
// @Summary Get Category Book  with Pagination
// @Description Get category Book with Pagination
// @Tags BookCategory
// @Param search query string false "search (string)"
// @Param page query integer false "page (int)"
// @Param limit query integer false "limit (int)"
// @Param status query integer false "status (status)"
// @Produce json
// @Success 200
// @Router /v1/book-category [get]
func GetBookCategories(c echo.Context) error {

	param := utils.PopulatePaging(c, "status")
	data := repository.GetBookCategories(param)

	return c.JSON(http.StatusOK, data)
}

// GetBookCategoryId godoc
// @Summary Get Single Book Category
// @Description Get Single Book Category
// @Tags BookCategory
// @Param any path string true "id"
// @Produce json
// @Success 200
// @Router /v1/book-category/{id} [get]

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

// DeleteBookCategoryId godoc
// @Summary Delete Single Book Category by ID
// @Description Delete Single Book Category by ID
// @Tags BookCategory
// @Produce json
// @Param id path integer true "id"
// @Success 200
// @Router /v1/book-category/{id} [delete]
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

// UpdateBookCategoryById godoc
// @Summary Update Single Book Category by ID
// @Description Update Single Book Category by ID
// @Tags BookCategory
// @Produce json
// @Param id path integer true "id"
// @Param Body body reqres.BookCategoryRequest true "Update body"
// @Success 200
// @Router /v1/book-category/{id} [put]
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

	if input.Icon != "" {
		data.Icon = input.Icon
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
