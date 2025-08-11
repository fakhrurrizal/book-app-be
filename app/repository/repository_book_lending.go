package repository

import (
	"book-app/app/models"
	"book-app/app/reqres"
	"book-app/app/utils"
	"book-app/config"
	"time"

	"github.com/guregu/null"
	"github.com/lib/pq"
)

func CreateBookLending(BorrowDate, DueDate, ReturnDate null.Time, data *reqres.BookLendingRequest) (response models.BookLending, err error) {

	response = models.BookLending{
		BookID:     data.BookID,
		UserID:     data.UserID,
		BorrowDate: BorrowDate,
		DueDate:    DueDate,
		ReturnDate: ReturnDate,
		Notes:      data.Notes,
		Status:     data.Status,
	}

	var created bool
	for !created {
		err = config.DB.Create(&response).Error
		if err != nil {
			if !config.LoadConfig().EnableIDDuplicationHandling {
				return
			}
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code != "23505" {
					return
				}
			}
		} else {
			created = true
		}
	}

	return
}

func GetBookLendings(bookID, userID int, param reqres.ReqPaging) (data reqres.ResPaging) {
	var responses []models.BookLending
	where := "book_lendings.deleted_at IS NULL"

	var modelTotal []models.BookLending

	type TotalResult struct {
		Total       int64
		LastUpdated time.Time
	}
	var totalResult TotalResult
	config.DB.Model(&modelTotal).Select("COUNT(*) AS total, MAX(updated_at) AS last_updated").Scan(&totalResult)

	// Mulai query
	query := config.DB.Table("book_lendings").
		Joins("JOIN books ON books.id = book_lendings.book_id").
		Joins("JOIN global_users ON global_users.id = book_lendings.user_id").
		Where(where)

	if param.Custom != "" {
		query = query.Where("book_lendings.status = ?", param.Custom)
	}
	if bookID > 0 {
		query = query.Where("book_lendings.book_id = ?", bookID)
	}

	if userID > 0 {
		query = query.Where("book_lendings.user_id = ?", userID)
	}

	if param.Search != "" {
		search := "%" + param.Search + "%"
		query = query.Where(`
        books.title ILIKE ? 
        OR global_users.fullname ILIKE ? 
        OR global_users.email ILIKE ?`,
			search, search, search,
		)
	}

	var totalFiltered int64
	query.Count(&totalFiltered)

	query.
		Select("book_lendings.*").
		Order(param.Sort + " " + param.Order).
		Limit(param.Limit).
		Offset(param.Offset).
		Scan(&responses)

	var responsesRefined []reqres.BookLendingResponse
	for _, item := range responses {
		responseRefined := BuildBookLendingResponse(item)
		responsesRefined = append(responsesRefined, responseRefined)
	}

	data = utils.PopulateResPaging(&param, responsesRefined, totalResult.Total, totalFiltered)
	return
}

func GetBookLendingByID(id int) (responseRefined reqres.BookLendingResponse, err error) {
	var response models.BookLending
	err = config.DB.First(&response, id).Error

	responseRefined = BuildBookLendingResponse(response)

	return
}

func GetBookLendingByIDPlain(id int) (response models.BookLending, err error) {
	err = config.DB.First(&response, id).Error

	return
}

func UpdateBookLending(request models.BookLending) (response models.BookLending, err error) {
	err = config.DB.Save(&request).Scan(&response).Error

	return
}

func DeleteBookLending(request models.BookLending) (models.BookLending, error) {
	err := config.DB.Delete(&request).Error
	return request, err
}

func BuildBookLendingResponse(data models.BookLending) (response reqres.BookLendingResponse) {

	var book models.Book
	var user models.GlobalUser

	response.CustomGormModel = data.CustomGormModel
	response.BorrowDate = data.BorrowDate
	response.DueDate = data.DueDate
	response.ReturnDate = data.ReturnDate
	response.Status = data.Status
	response.Notes = data.Notes

	if data.BookID > 0 {
		book, _ = GetBookByIDPlain(data.BookID)
		response.Book = reqres.GlobalIDNameResponse{
			ID:    int(book.ID),
			Name:  book.Title,
			Image: book.Image,
		}
	}

	if data.UserID > 0 {
		user, _ = GetUserByIDPlain(data.UserID)
		response.User = reqres.GlobalIDNameResponse{
			ID:    int(user.ID),
			Name:  user.Fullname,
			Email: user.Email,
		}
	}

	return response
}
