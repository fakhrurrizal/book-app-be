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

func CreateBookCategory(data *reqres.BookCategoryRequest) (response models.BookCategory, err error) {
	response = models.BookCategory{
		Name:        data.Name,
		Description: data.Description,
		Status:      data.Status,
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

func BuildBookCategoryResponse(data models.BookCategory) (response reqres.BookCategoryResponse) {

	response.CustomGormModel = reqres.CustomGormModel(data.CustomGormModel)
	response.Name = data.Name
	response.Description = data.Description
	response.Status = data.Status
	response.Icon = data.Icon

	return
}

func GetAllBookCategory(search string, categoryID int) ([]models.BookCategory, error) {
	var responses []models.BookCategory

	where := "deleted_at IS NULL AND status = true"

	err := config.DB.Find(&responses).Error
	if err != nil {
		return nil, err
	}

	if categoryID > 0 {
		where += " AND category_id = " + strconv.Itoa(categoryID)
	}

	if search != "" {
		var filteredCategories []models.BookCategory
		for _, category := range responses {
			if strings.Contains(strings.ToLower(category.Name), strings.ToLower(search)) ||
				strings.Contains(strings.ToLower(category.Description), strings.ToLower(search)) {
				filteredCategories = append(filteredCategories, category)
			}
		}
		return filteredCategories, nil
	}

	return responses, nil
}

func GetBookCategories(param reqres.ReqPaging) (data reqres.ResPaging) {
	var responses []models.BookCategory

	where := "deleted_at IS NULL"

	var modelTotal []models.BookCategory

	var totalResult int64
	config.DB.Debug().Model(&modelTotal).Where(where).Count(&totalResult)

	var totalFiltered int64

	config.DB.Limit(param.Limit).Offset(param.Offset).Order(param.Sort + " " + param.Order).Where(where).Find(&responses)

	var responsesRefined []reqres.BookCategoryResponse
	if param.Search != "" {
		for _, item := range responses {
			if strings.Contains(strings.ToLower(item.Name), strings.ToLower(param.Search)) {
				responseRefined := BuildBookCategoryResponse(item)

				responsesRefined = append(responsesRefined, responseRefined)
			}
		}
	} else {
		for _, item := range responses {
			responseRefined := BuildBookCategoryResponse(item)

			responsesRefined = append(responsesRefined, responseRefined)
		}
	}

	data = utils.PopulateResPaging(&param, responsesRefined, totalResult, totalFiltered)

	return
}

func GetBookCategoryID(id int) (response reqres.BookCategoryResponse, err error) {
	var data models.BookCategory
	err = config.DB.First(&data, id).Error

	response = BuildBookCategoryResponse(data)

	return
}

func GetBookCategoryIDPlain(id int) (response models.BookCategory, err error) {
	err = config.DB.First(&response, id).Error

	return
}

func DeleteBookCategory(request models.BookCategory) (models.BookCategory, error) {
	err := config.DB.Delete(&request).Error

	return request, err
}

func UpdateBookCategory(request models.BookCategory) (response models.BookCategory, err error) {
	err = config.DB.Save(&request).Scan(&response).Error

	return
}
