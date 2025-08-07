package repository

import (
	"book-app/app/models"
	"book-app/app/reqres"
	"book-app/app/utils"
	"book-app/config"
	"strconv"
	"time"

	"github.com/lib/pq"
)

func CreateRole(appID int, data *reqres.GlobalRoleRequest) (response models.GlobalRole, err error) {

	response = models.GlobalRole{
		Name:        data.Name,
		Description: data.Description,
		Status:      data.Status,
		AppID:       appID,
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

func GetRoles(appID int, param reqres.ReqPaging) (data reqres.ResPaging) {
	var responses []models.GlobalRole
	where := "deleted_at IS NULL"

	if appID > 0 {
		where += " AND app_id = " + strconv.Itoa(appID)
	}

	var modelTotal []models.GlobalRole

	type TotalResult struct {
		Total       int64
		LastUpdated time.Time
	}
	var totalResult TotalResult
	config.DB.Model(&modelTotal).Select("COUNT(*) AS total, MAX(updated_at) AS last_updated").Scan(&totalResult)

	if param.Custom != "" {
		where += " AND status = " + param.Custom.(string)
	}
	if param.Search != "" {
		where += " AND name ILIKE '%" + param.Search + "%'"
	}

	var totalFiltered int64
	config.DB.Model(&modelTotal).Where(where).Count(&totalFiltered)

	config.DB.Limit(param.Limit).Offset(param.Offset).Order(param.Sort + " " + param.Order).Where(where).Find(&responses)

	var responsesRefined []reqres.GlobalRoleResponse
	for _, item := range responses {
		responseRefined := BuildRoleResponse(item)

		responsesRefined = append(responsesRefined, responseRefined)
	}

	data = utils.PopulateResPaging(&param, responsesRefined, totalResult.Total, totalFiltered)

	return
}

func GetRoleByID(id int) (responseRefined reqres.GlobalRoleResponse, err error) {
	var response models.GlobalRole
	err = config.DB.First(&response, id).Error

	responseRefined = BuildRoleResponse(response)

	return
}

func GetRoleByIDPlain(id int) (response models.GlobalRole, err error) {
	err = config.DB.First(&response, id).Error

	return
}

func UpdateRole(request models.GlobalRole) (response models.GlobalRole, err error) {
	err = config.DB.Save(&request).Scan(&response).Error

	return
}

func DeleteRole(request models.GlobalRole) (models.GlobalRole, error) {
	err := config.DB.Delete(&request).Error
	return request, err
}

func BuildRoleResponse(data models.GlobalRole) (response reqres.GlobalRoleResponse) {

	response.CustomGormModel = data.CustomGormModel
	response.Name = data.Name
	response.Status = data.Status
	response.Description = data.Description

	return response
}
