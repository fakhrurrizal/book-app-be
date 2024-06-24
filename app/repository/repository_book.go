package repository

import (
	"book-app/app/models"
	"book-app/app/reqres"
	"book-app/app/utils"
	"book-app/config"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

func CreateBook(data *reqres.BookRequest) (response models.Book, err error) {

	response = models.Book{
		Title:           data.Title,
		Author:          data.Author,
		PublicationYear: data.PublicationYear,
		Language:        data.Language,
		Description:     data.Description,
		Publisher:       data.Publisher,
		NumberOfPages:   data.NumberOfPages,
		BookCode:        data.BookCode,
		CategoryID:      data.CategoryID,
		Image:           data.Image,
		Status:          data.Status,
	}

	var created bool
	for !created {
		err = config.DB.Create(&response).Error
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code != "23505" {
					created = true
				}
			}
		} else {
			created = true
		}
	}

	return
}

func BuildBookResponse(data models.Book) (response reqres.BookResponse) {

	books := reqres.BookResponse{
		CustomGormModel: data.CustomGormModel,
		Title:           data.Title,
		Author:          data.Author,
		Image:           data.Image,
		PublicationYear: data.PublicationYear,
		Language:        data.Language,
		BookCode:        data.BookCode,
		Publisher:       data.Publisher,
		Description:     data.Description,
		NumberOfPages:   data.NumberOfPages,
		Status:          data.Status,
	}
	if data.CategoryID > 0 {
		category, _ := GetBookCategoryIDPlain(data.CategoryID)
		books.Category.ID = int(category.ID)
		books.Category.Name = category.Name
	}

	response = books

	return response

}

func GetBookByIDPlain(id int) (response models.Book, err error) {
	err = config.DB.First(&response, id).Error

	return
}

func GetBookByID(id int) (responseRefined reqres.BookResponse, err error) {

	var response models.Book

	err = config.DB.First(&response, id).Error

	responseRefined = BuildBookResponse(response)

	return
}

func GetBooks(categoryID int, param reqres.ReqPaging) (data reqres.ResPaging) {
	var responses []models.Book

	where := "books.deleted_at IS NULL"

	if categoryID > 0 {
		where += " AND books.category_id = " + strconv.Itoa(categoryID)
	}

	if param.Custom != "" {
		where += " AND books.status = " + param.Custom.(string)
	}

	var totalResult int64
	config.DB.Debug().Model(&models.Book{}).Where(where).Count(&totalResult)

	config.DB.Debug().Where(where).Find(&responses)

	bookQueue := &utils.BookQueue{}
	if param.Search != "" {
		for _, book := range responses {
			if strings.Contains(strings.ToLower(book.Title), strings.ToLower(param.Search)) ||
				strings.Contains(strings.ToLower(book.Description), strings.ToLower(param.Search)) {
				bookQueue.Enqueue(BuildBookResponse(book))
			}
		}
	} else {
		for _, book := range responses {
			bookQueue.Enqueue(BuildBookResponse(book))
		}
	}

	var filteredResponses []reqres.BookResponse
	for !bookQueue.IsEmpty() {
		book, _ := bookQueue.Dequeue()
		filteredResponses = append(filteredResponses, book)
	}

	totalFiltered := int64(len(filteredResponses))

	sortBooks(&filteredResponses, param.Sort, param.Order)

	startIndex := param.Offset
	endIndex := param.Offset + param.Limit
	if endIndex > len(filteredResponses) {
		endIndex = len(filteredResponses)
	}
	pagedResponses := filteredResponses[startIndex:endIndex]

	data = utils.PopulateResPaging(&param, pagedResponses, totalResult, totalFiltered)

	return data
}

func sortBooks(books *[]reqres.BookResponse, sortField string, sortOrder string) {
	switch sortField {
	case "id":
		bubbleSort(books, func(a, b reqres.BookResponse) bool {
			if sortOrder == "ASC" {
				return a.ID < b.ID
			}
			return a.ID > b.ID
		})
	case "publication_year":
		bubbleSort(books, func(a, b reqres.BookResponse) bool {
			if sortOrder == "ASC" {
				return a.PublicationYear < b.PublicationYear
			}
			return a.PublicationYear > b.PublicationYear
		})
	default:
		bubbleSort(books, func(a, b reqres.BookResponse) bool {
			return a.ID < b.ID
		})
	}
}

func GetAllBooksPlain(search string, categoryID int) (responses []models.Book, err error) {
	where := "deleted_at IS NULL AND status = true"

	if search != "" {
		where += " AND name ILIKE '%" + search + "%' OR code ILIKE '%" + search + "%' AND barcode ILIKE '%" + search + "%'"
	}

	if categoryID > 0 {
		where += " AND category_id = " + strconv.Itoa(categoryID)
	}

	err = config.DB.Where(where).Find(&responses).Error

	return
}

func bubbleSort(books *[]reqres.BookResponse, less func(a, b reqres.BookResponse) bool) {
	n := len(*books)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if !less((*books)[j], (*books)[j+1]) {
				(*books)[j], (*books)[j+1] = (*books)[j+1], (*books)[j]
			}
		}
	}
}

func DeleteBook(request models.Book) (models.Book, error) {

	err := config.DB.Delete(&request).Error

	return request, err
}

func UpdateBook(request models.Book) (response models.Book, err error) {
	err = config.DB.Save(&request).Scan(&response).Error

	return
}
