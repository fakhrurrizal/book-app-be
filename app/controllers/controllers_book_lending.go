package controllers

import (
	"book-app/app/repository"
	"book-app/app/reqres"
	"book-app/app/utils"
	"net/http"
	"strconv"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/guregu/null"
	"github.com/labstack/echo/v4"
)

// CreateBookLending godoc
// @Summary Create BookLending
// @Description Create New BookLending
// @Tags BookLending
// @Produce json
// @Param Body body reqres.BookLendingRequest true "Create body"
// @Success 200
// @Router /v1/book-lending [post]
// @Security ApiKeyAuth
// @Security JwtToken
func CreateBookLending(c echo.Context) error {
	var input reqres.BookLendingRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}
	utils.StripTagsFromStruct(&input)

	if err := input.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(400, utils.NewInvalidInputError(errVal))
	}

	var borrowDate time.Time
	var dueDate time.Time
	var returnDate time.Time
	var err error

	if input.BorrowDate != "" {
		borrowDate, err = time.ParseInLocation("2006-01-02 15:04:05", input.BorrowDate, utils.GetTimeLocation())
		if err != nil {
			borrowDate, err = time.ParseInLocation(time.DateOnly, input.BorrowDate, utils.GetTimeLocation())
			if err != nil {
				return c.JSON(500, utils.Respond(500, err, "Invalid Borrow date format"))
			}
		}
	}

	if input.DueDate != "" {
		dueDate, err = time.ParseInLocation("2006-01-02 15:04:05", input.DueDate, utils.GetTimeLocation())
		if err != nil {
			dueDate, err = time.ParseInLocation(time.DateOnly, input.DueDate, utils.GetTimeLocation())
			if err != nil {
				return c.JSON(500, utils.Respond(500, err, "Invalid Due date format"))
			}
		}
	}

	if input.ReturnDate != "" {
		returnDate, err = time.ParseInLocation("2006-01-02 15:04:05", input.ReturnDate, utils.GetTimeLocation())
		if err != nil {
			returnDate, err = time.ParseInLocation(time.DateOnly, input.ReturnDate, utils.GetTimeLocation())
			if err != nil {
				return c.JSON(500, utils.Respond(500, err, "Invalid Return date format"))
			}
		}
	}

	data, err := repository.CreateBookLending(null.TimeFrom(borrowDate), null.TimeFrom(dueDate), null.TimeFrom(returnDate), &input)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to create"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "Create Success",
	})
}

// GetBookLendings godoc
// @Summary Get All BookLending With Pagination
// @Description Get All BookLending With Pagination
// @Tags BookLending
// @Param search query string false "search (string)"
// @Param page query integer false "page (int)"
// @Param limit query integer false "limit (int)"
// @Param sort query string false "sort (ASC/DESC)"
// @Param order query string false "order by (default: id)"
// @Param status query string false "status ('requested','approved', 'borrowed','rejected' 'returned', 'overdue')"
// @Param book_id query integer false "book_id (int)"
// @Param user_id query integer false "user_id (int)"
// @Param code query string false "code (string)"
// @Produce json
// @Success 200
// @Router /v1/book-lending [get]
// @Security ApiKeyAuth
func GetBookLendings(c echo.Context) error {
	bookId, _ := strconv.Atoi(c.QueryParam("book_id"))
	userId, _ := strconv.Atoi(c.QueryParam("user_id"))
	param := utils.PopulatePaging(c, "status")
	data := repository.GetBookLendings(bookId, userId, param)

	return c.JSON(http.StatusOK, data)
}

// GetBookLendingByID godoc
// @Summary Get Single BookLending
// @Description Get Single BookLending
// @Tags BookLending
// @Param id path integer true "ID"
// @Produce json
// @Success 200
// @Router /v1/book-lending/{id} [get]
// @Security ApiKeyAuth
func GetBookLendingByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repository.GetBookLendingByID(id)
	if err != nil {
		return c.JSON(404, utils.Respond(404, err, "Failed to get"))
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "Success to get",
	})
}

// DeleteBookLendingByID godoc
// @Summary Delete Single BookLending by ID
// @Description Delete Single BookLending by ID
// @Tags BookLending
// @Produce json
// @Param id path integer true "ID"
// @Success 200
// @Router /v1/book-lending/{id} [delete]
// @Security ApiKeyAuth
// @Security JwtToken
func DeleteBookLendingByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repository.GetBookLendingByIDPlain(id)
	if err != nil {
		return c.JSON(500, utils.Respond(404, err, "Failed to get"))
	}

	_, err = repository.DeleteBookLending(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to delete"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "Success to delete",
	})
}

// UpdateBookLendingByID godoc
// @Summary Update Single BookLending by ID
// @Description Update Single BookLending by ID
// @Tags BookLending
// @Produce json
// @Param id path integer true "ID"
// @Param Body body reqres.BookLendingRequest true "Update body"
// @Success 200
// @Router /v1/book-lending/{id} [put]
// @Security ApiKeyAuth
// @Security JwtToken
func UpdateBookLendingByID(c echo.Context) error {
	var input reqres.BookLendingRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}
	utils.StripTagsFromStruct(&input)

	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repository.GetBookLendingByIDPlain(id)
	if err != nil {
		return c.JSON(500, utils.Respond(404, err, "Failed to get"))
	}

	if input.BookID != 0 {
		data.BookID = input.BookID
	}

	if input.UserID != 0 {
		data.UserID = input.UserID
	}

	if input.Status != "" {
		data.Status = input.Status
	}

	data.Notes = input.Notes

	data.Status = input.Status

	if input.BorrowDate != "" {
		borrow_date, err := time.ParseInLocation("2006-01-02 15:04:05", input.BorrowDate, utils.GetTimeLocation())
		if err != nil {
			borrow_date, err = time.ParseInLocation(time.DateOnly, input.BorrowDate, utils.GetTimeLocation())
			if err != nil {
				return c.JSON(400, utils.Respond(400, err, "Invalid borrow_date date format"))
			}
		}
		data.BorrowDate = null.TimeFrom(borrow_date)
	}

	if input.DueDate != "" {
		due_date, err := time.ParseInLocation("2006-01-02 15:04:05", input.DueDate, utils.GetTimeLocation())
		if err != nil {
			due_date, err = time.ParseInLocation(time.DateOnly, input.DueDate, utils.GetTimeLocation())
			if err != nil {
				return c.JSON(400, utils.Respond(400, err, "Invalid due_date date format"))
			}
		}
		data.DueDate = null.TimeFrom(due_date)
	}

	if input.ReturnDate != "" {
		return_date, err := time.ParseInLocation("2006-01-02 15:04:05", input.ReturnDate, utils.GetTimeLocation())
		if err != nil {
			return_date, err = time.ParseInLocation(time.DateOnly, input.ReturnDate, utils.GetTimeLocation())
			if err != nil {
				return c.JSON(400, utils.Respond(400, err, "Invalid return_date date format"))
			}
		}
		data.ReturnDate = null.TimeFrom(return_date)
	}

	dataUpdate, err := repository.UpdateBookLending(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to update"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    dataUpdate,
		"message": "Success to update",
	})
}
