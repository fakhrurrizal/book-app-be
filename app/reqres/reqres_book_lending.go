package reqres

import (
	"book-app/app/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/guregu/null"
)

type BookLendingRequest struct {
	BookID     int    `json:"book_id"`
	UserID     int    `json:"user_id"`
	BorrowDate string `json:"borrow_date"`
	DueDate    string `json:"due_date"`
	ReturnDate string `json:"return_date"`
	Status     string `json:"status"`
	Notes      string `json:"notes"`
}

func (request BookLendingRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.BookID, validation.Required),
		validation.Field(&request.UserID, validation.Required),
		validation.Field(&request.Status, validation.Required),
	)
}

type BookLendingResponse struct {
	models.CustomGormModel
	Book       GlobalIDNameResponse `json:"book"`
	User       GlobalIDNameResponse `json:"user"`
	BorrowDate null.Time            `json:"borrow_date"`
	DueDate    null.Time            `json:"due_date"`
	ReturnDate null.Time            `json:"return_date"`
	Status     string               `json:"status"`
	Notes      string               `json:"notes"`
}
